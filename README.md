# Loan Service
### Description
This project is a loan service platform designed to manage loans, user actions, investments, and send automated email agreements to investors. The service is built with Go and uses various components such as NSQ for message queueing, GORM for database interactions, and SMTP (Gmail) for sending email notifications.

### Assumptions
- One user only has one role: A user can only have one role at a time, such as admin, borrower, or investor.
- One borrower can only have one loan at a time: Each borrower is restricted to a single loan.
- Source of funds is not using a third party: The service manages funds internally without relying on third-party services.
- Action flow is sequential: A loan must go through specific stages (e.g., approval, disbursement) in a sequential order.
- User information is unique: Each user has a unique email address, which is used for communication purposes (e.g., sending agreement letters).

### Prerequisites
Before setting up the project, make sure you have the following installed on your machine:
```
Go 1.18+
PostgreSQL
NSQ
Gmail Account for SMTP Email
wkhtmltopdf for generating PDFs (if needed)
```

### Setup Instructions
1. Clone the Repository
Start by cloning the repository to your local machine:

```
git clone https://github.com/raihaniyai/loan-service.git
cd loan-service
```

2. Set Up Your Environment Variables
You need to configure the following environment variables. These will be used to connect to the database, NSQ, and SMTP service:

export configs value:
```
export DB_USER=your_db_user
export DB_PASSWORD=your_db_password
export DB_NAME=loan_service_db
export DB_HOST=localhost
export DB_PORT=5432

export NSQD_ADDRESS=localhost:4150
export NSQ_LOOKUPD_ADDRESS=localhost:4161
export NSQ_TOPIC=loan-investment-completed
export NSQ_CHANNEL=loan-channel

export SMTP_HOST=smtp.gmail.com
export SMTP_EMAIL=your-email@gmail.com
export SMTP_PASSWORD=your-app-password
```

Replace the placeholder values with your actual configurations:

For DB_USER, DB_PASSWORD, and other DB-related fields, use your PostgreSQL credentials.
For NSQD_ADDRESS and NSQ_LOOKUPD_ADDRESS, use your NSQ service configuration.
For SMTP_HOST, SMTP_EMAIL, and SMTP_PASSWORD, use your Gmail credentials and app password.

3. Install Dependencies
Make sure you have Go and all necessary dependencies installed. From the root of the project, run:

```
go mod tidy
```

This will download all the required dependencies.

4. Set Up the Database
Run the SQL DDL (Data Definition Language) commands to set up the necessary database tables. You can execute the following SQL DDL commands in your PostgreSQL database to create the required tables:

```
CREATE TABLE loans (
    loan_id BIGSERIAL PRIMARY KEY,
    borrower_id BIGINT NOT NULL,
    principal_amount BIGINT NOT NULL,
    interest_rate REAL NOT NULL,        -- Stored as percentage
    return_on_investment REAL NOT NULL, -- Stored as percentage
    agreement_letter TEXT,
    status INT DEFAULT 10,
    updated_by BIGINT,                  -- Stores user ID, assuming from external service
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

-- Indexes for loans table
CREATE INDEX idx_loans_borrower_id ON loans(borrower_id);
CREATE INDEX idx_loans_status ON loans(status);
CREATE INDEX idx_loans_updated_by ON loans(updated_by);
CREATE INDEX idx_loans_created_at ON loans(created_at);

CREATE TABLE actions (
    action_id BIGSERIAL PRIMARY KEY,
    loan_id BIGINT NOT NULL REFERENCES loans(loan_id) ON DELETE CASCADE,
    action_type INT NOT NULL,           -- Enum or predefined action types (e.g., approval, disbursement)
    action_by INT NOT NULL,
    document_url TEXT,
    created_by BIGINT NOT NULL,         -- Stores user ID, assuming from external service
    created_at TIMESTAMP NOT NULL
);

-- Indexes for actions table
CREATE INDEX idx_actions_loan_id ON actions(loan_id);
CREATE INDEX idx_actions_action_type ON actions(action_type);
CREATE INDEX idx_actions_action_by ON actions(action_by);
CREATE INDEX idx_actions_created_at ON actions(created_at);

CREATE TABLE investments (
    investment_id BIGSERIAL PRIMARY KEY,
    loan_id BIGINT NOT NULL REFERENCES loans(loan_id) ON DELETE CASCADE,
    investor_id BIGINT NOT NULL,        -- Stores user ID, assuming from external service
    investment_amount BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL
);

-- Indexes for investments table
CREATE INDEX idx_investments_loan_id ON investments(loan_id);
CREATE INDEX idx_investments_investor_id ON investments(investor_id);
CREATE INDEX idx_investments_created_at ON investments(created_at);

CREATE TABLE users (
    user_id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255),
    role INT NOT NULL,                  -- Enum or predefined roles (e.g., admin=1, borrower=2, investor=3)
    email VARCHAR(255) UNIQUE,          -- contains PII, need to be masked
    phone_number VARCHAR(20)            -- contains PII, need to be masked
);

-- Indexes for users table
CREATE INDEX idx_users_role ON users(role);
CREATE INDEX idx_users_email ON users(email);

CREATE TABLE funds (
    fund_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,      -- Stores user ID from external service
    balance BIGINT NOT NULL,      -- Represents the available investment amount
    created_at TIMESTAMP NOT NULL 
);

-- Index for efficient lookup of funds by user ID
CREATE INDEX idx_funds_user_id ON funds(user_id);
```

Execute these commands in your PostgreSQL instance to set up the database.

5. Configure NSQ
If you haven't already, you need to install and run NSQ for message queueing.

Start the nsqd and nsqlookupd services.
```
nsqd --lookupd-tcp-address=localhost:4160
nsqlookupd
```
Make sure that NSQD_ADDRESS and NSQ_LOOKUPD_ADDRESS in your .env file match the addresses of your local NSQ services.

6. Run the Application
Once you have everything configured, run the application using:

```
make run
```
This will start the HTTP server and listen for incoming API requests.

### Postman Collection
You can import the provided Postman collection to test the API endpoints. The collection contains requests for the following:
```
POST /loans: Create a new loan.
POST /loans/{loanID}/approve: Approve a loan.
POST /loans/{loanID}/disburse: Disburse a loan.
POST /loans/{loanID}/invest: Invest in a loan.
POST /funds/topup: Top-up user balance.
POST /users: Create a new user.
```

Import Postman Collection
Download the Postman collection file.
Open Postman, go to the "Import" option, and select the downloaded file.
You can now test the API endpoints.

### Send Agreement Letter
The service automatically generates an agreement letter PDF for investors when an investment is completed. The PDF is sent to the investor's email address through SMTP.

### Trigger Process:
When an investment is completed (loan-investment-completed message is published to NSQ), the service listens for this event.
The service generates a PDF agreement letter and sends it to the investor's email via SMTP.
Make sure to check that the SMTP configurations are correct to ensure email sending works.

### Troubleshooting
Error: exec: "wkhtmltopdf": executable file not found in $PATH
Solution: Ensure that wkhtmltopdf is installed and available in your system’s $PATH. Download wkhtmltopdf.

Error: Failed to generate PDF: exit status 1
Solution: Check the output logs to identify potential issues with generating PDFs. Ensure that the HTML template is valid and that wkhtmltopdf is functioning properly.

Error: Unable to write to destination
Solution: Ensure that the file path where you are saving the PDF is writable by the user running the service.
