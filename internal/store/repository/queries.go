package repository

const createRepoQuery = `
	INSERT INTO repositories (name) VALUES ($1)
`
