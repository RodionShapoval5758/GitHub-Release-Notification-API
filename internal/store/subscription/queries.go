package subscription

const createSubscriptionQuery = `
	INSERT INTO subscriptions (
		email,
		repository_id,
		confirmation_token,
		unsubscribe_token
	)
	VALUES ($1, $2, $3, $4)
`

const findSubscriptionByEmailAndRepositoryIDQuery = `
	SELECT id, email, repository_id, confirmed, confirmation_token, unsubscribe_token, created_at, confirmed_at
	FROM subscriptions
	WHERE email = $1 AND repository_id = $2
`

const findSubscriptionByConfirmTokenQuery = `
	SELECT id, email, repository_id, confirmed, confirmation_token, unsubscribe_token, created_at, confirmed_at
	FROM subscriptions
	WHERE confirmation_token = $1
`

const findSubscriptionByUnsubscribeTokenQuery = `
	SELECT id, email, repository_id, confirmed, confirmation_token, unsubscribe_token, created_at, confirmed_at
	FROM subscriptions
	WHERE unsubscribe_token = $1
`

const confirmSubscriptionByTokenQuery = `
	UPDATE subscriptions
	SET confirmed = TRUE, confirmed_at = NOW()
	WHERE confirmation_token = $1
`

const deleteSubscriptionByUnsubscribeTokenQuery = `
	DELETE FROM subscriptions
	WHERE unsubscribe_token = $1
`

const listConfirmedSubscriptionsByEmailQuery = `
	SELECT id, email, repository_id, confirmed, confirmation_token, unsubscribe_token, created_at, confirmed_at
	FROM subscriptions
	WHERE email = $1 AND confirmed = TRUE
`
