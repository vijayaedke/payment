package utils

import "errors"

type ErrorCodes string

const (
	INVALID_ACC_REQUEST_PAYLOAD ErrorCodes = "ACC_ERR_001"
	CREATE_ACC_FAILED           ErrorCodes = "ACC_ERR_002"

	ACC_NOT_FOUND        ErrorCodes = "ACC_ERR_003"
	INVALID_ACC_ID       ErrorCodes = "ACC_ERR_004"
	FAILED_ACC_ID_DETAIL ErrorCodes = "ACC_ERR_005"

	INVALID_TXN_REQUEST_PAYLOAD ErrorCodes = "TXN_ERR_001"
	FAILED_TXN_CREATE           ErrorCodes = "TXN_ERR_002"
	OP_ID_NOT_FOUND             ErrorCodes = "TXN_ERR_003"
)

type ErrorMsg string

const (
	CREATE_ACC_FAILED_MSG           ErrorMsg = "Failed account creation"
	ACC_NOT_FOUND_MSG               ErrorMsg = "Account details not found/doesn't exists"
	INVALID_ACC_REQUEST_PAYLOAD_MSG ErrorMsg = "Invalid account creation payload"
	INVALID_ACC_ID_MSG              ErrorMsg = "Invalid account id format provided"
	INVALID_TXN_REQUEST_PAYLOAD_MSG ErrorMsg = "Invalid transaction request payload"
	FAILED_TXN_CREATE_MSG           ErrorMsg = "Failed to create transaction details"
	OP_ID_NOT_FOUND_MSG             ErrorMsg = "Operation id details not found/doesn't exists"
	FAILED_ACC_ID_DETAIL_MSG        ErrorMsg = "Failed to get account id details"
)

var AccIdNotExists = errors.New("account id not found")
var OpIdNotExists = errors.New("operation id not found")

const ACCOUNT_ID_PARAM = "account_id"
const OPERTAION_ID_PARAM = "operation_type_id"
