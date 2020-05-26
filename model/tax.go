package model

type account struct {
	account_num      string
	account_currency string
}

type request struct {
	inn                  string
	bank_branch_inn      string
	doc_date             string
	doc_num              string
	bank_branch_director string
	bank_branch_operator string
	authCode             string
	accounts             []account
}
