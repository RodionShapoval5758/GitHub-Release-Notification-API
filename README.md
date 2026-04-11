# 
GitHub-Release-Notification-API
API that allows users to subscribe to email notifications about new releases of a chosen GitHub repository.


Workflow + thoughts:

DB schema:
    Two main entities:
        subscription(email, repo id, tokens, confirmed), 
        repo(name, last_checked)
    subscriptions table has a composite UNIQUE that ensures the same pair of email and repo cannot be added 
    So that the last changes are not updated in all the users separately, but in one row
    + indexes for tokens and email

Errors && logging:
    Wrapping technical failures(db query failed, network failure)
    Sentinel errors for meaningful errors (not found, exists, invalid format)
    Logging:
        ERRORS: log fatal with context when something failed and needs attention
        INFO: when some process started or for providing context, so that I know what has happened

Instead of creating GerOrCreate method for repository that is race-safe
I separated methods Create and Find and implemented manual race condition handling
so that everything is explicit and content

In case of duplicate tokens the program tries to regenerate it 5 times. 
There is a ridiculously small possibility of getting 
duplicate tokens with 32 token length, so I neglect it

Swagger docs states that /api/subscribe consumes both json and form data, so I made a fallback for both