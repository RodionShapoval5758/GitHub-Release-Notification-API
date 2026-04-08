# 
GitHub-Release-Notification-API
API that allows users to subscribe to email notifications about new releases of a chosen GitHub repository.


Workflow + thoughts:

DB schema:
    Two main entities:
        subscription(email, repo id, tokens, confirmed), 
        repo(name, last_checked)
    So that the last changes are not updated in all the users separately, but in one row
    + indexes for tokens and email

Errors && logging:
    Wrapping technical failures(db query failed, network failure)
    Sentinel errors for meaningful errors (not found, exists, invalid format)
    Logging:
        ERRORS: log fatal with context when something failed and needs attention
        INFO: when some process started or for providing context, so that I know what has happened