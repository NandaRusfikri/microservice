package constant

const (
	PREFIX_PROJECT             = "SDH"
	APPLICATION_NAME           = "WaterMeter"
	KEY_PASSWORD               = "V1si0n3t*1!"
	KEY_ENCRYPT_AES            = "KitaCobaCekDulu!"
	IV_ENCRYPT_AES             = "KitaCobaCekDulu!"
	LENGTH_TOKEN_DEFAULT       = 16
	LONG_TOKEN_EXPIRED_DEFAULT = 1440
	JWT_SIGNATURE_KEY          = "VSN-EDC-2022"

	PAGING_SYNC_DEFAULT = 1
	LIMIT_SYNC_DEFAULT  = 20

	STORAGE_PATH_PAYMENT     = "./storage/payment"
	STORAGE_PATH_BANNNER     = "./storage/banner"
	STORAGE_PATH_MERCHANT    = "./storage/merchant"
	STORAGE_PATH_TASKLIST    = "./storage/tasklist"
	STORAGE_PATH_APPLICATION = "./storage/application"

	PAYMENT_ID_ATOME   = 1
	PAYMENT_ID_MANAGER = 2
	BANK_PARTNER_NOBU  = 3

	STATUS_ACTIVE   = "ACTIVE"
	STATUS_INACTIVE = "INACTIVE"

	ORDER_STATUS_PENDING = "PENDING"
	ORDER_STATUS_FAILED  = "FAILED"
	ORDER_STATUS_SUCCESS = "SUCCESS"
	ORDER_STATUS_CANCEL  = "CANCELLED"
	ORDER_STATUS_REFUND  = "REFUNDED"

	ORDER_MANAGER_STATUS_SETTLEMENT     = "SETTLEMENT"
	ORDER_MANAGER_STATUS_NOT_SETTLEMENT = "NOT SETTLEMENT"
	ORDER_MANAGER_STATUS_SALE           = "SALE"
	ORDER_MANAGER_STATUS_VOID           = "VOID"
	ORDER_MANAGER_STATUS_FAILED         = "FAILED"

	STATUS_OPEN           = 1
	STATUS_SUBMIT           = 2
	STATUS_DONE           = 3
	STATUS_RESCHEDULE           = 4

	ATOME_PAYMENT_PAID       = "PAID"
	ATOME_PAYMENT_CANCELLED  = "CANCELLED"
	ATOME_PAYMENT_REFUNDED   = "REFUNDED"
	ATOME_PAYMENT_PROCESSING = "PROCESSING"
	ATOME_PAYMENT_FAILED     = "FAILED"

	TABLE_PAYMENT_CHANNEL_ATOME = "ATOME"

	TABLE_MERCHANT_CHANNEL_NAME         = "merchant_channel"
	TABLE_MERCHANT_CHANNEL_NAME_PRELOAD = "MerchantChannel"

	TABLE_REASON_NAME                = "reason"
	TABLE_STATUS_NAME                = "status"
	TABLE_STATUS_NAME_PRELOAD   = "Status"
	TABLE_BIN_TMS_NAME               = "bin_tms"
	TABLE_REASON_NAME_PRELOAD        = "Reason"
	TABLE_PAYMENT_NAME               = "payment"
	TABLE_PAYMENT_NAME_PRELOAD       = "Payment"
	TABLE_TERMINAL_NAME              = "terminal"
	TABLE_TERMINAL_NAME_PRELOAD      = "Terminal"
	TABLE_BANK_ACQUIRED_NAME_PRELOAD = "BankAcquired"
	TABLE_TERMINAL_lOG_NAME          = "terminal_log"
	TABLE_TERMINAL_LOG_NAME_PRELOAD  = "TerminalLog"

	TABLE_BANNERS_NAME         = "banners"
	TABLE_BANNERS_NAME_PRELOAD = "Banners"

	TABLE_PROJECT_NAME                = "project"
	TABLE_PROJECT_NAME_PRELOAD        = "Project"
	TABLE_UNIT_NAME               = "unit"
	TABLE_TASKLIST_NAME               = "tasklist"
	TABLE_MERCHANT_NAME               = "tasklist"
	TABLE_TASKLIST_NAME_PRELOAD          = "Tasklist"
	TABLE_TASKLIST_UNIT_NAME               = "tasklist_unit"
	TABLE_TASKLIST_UNIT_NAME_PRELOAD           = "TasklistSurveys"
	TABLE_MERCHANT_GROUP_NAME         = "merchant_group"
	TABLE_MERCHANT_GROUP_NAME_PRELOAD = "MerchantGroup"
	TABLE_MERCHANT_NAME_PRELOAD       = "Merchant"
	TABLE_FEATURE_NAME                = "feature"
	TABLE_FEATURE_NAME_PRELOAD        = "Feature"

	TABLE_ROLE_NAME            = "role"
	TABLE_ROLE_PERMISSION_NAME = "role_permission"

	TABLE_USER_NAME                    = "users"
	TABLE_USER_LOG_NAME                = "user_log"
	TABLE_USER_ACTIVITY_LOG_NAME       = "user_activity_log"
	TABLE_USER_NAME_PRELOAD            = "User"
	TABLE_USER_CREDENTIAL_NAME         = "user_credential"
	TABLE_USER_CREDENTIAL_NAME_PRELOAD = "UserCredential"
	TABLE_USER_ROLE_NAME               = "user_role"
	TABLE_USER_ROLE_NAME_PRELOAD       = "UserRole"
	TABLE_USER_MERCHANT_NAME           = "user_merchant"
	TABLE_USER_MERCHANT_NAME_PRELOAD   = "UserMerchant"

	TABLE_OTP_NAME = "otp"

	TABLE_VERSION_NAME = "version"

	TABLE_CONFIG_NAME = "config"

	TABLE_ORDER_NAME             = "orders"
	TABLE_ORDER_NAME_PRELOAD     = "Orders"
	TABLE_ORDER_LOG_NAME         = "order_log"
	TABLE_ORDER_LOG_NAME_PRELOAD = "OrderLog"

	OPERATION_SQL_INSERT = "insert"
	OPERATION_SQL_UPDATE = "update"
	OPERATION_SQL_DELETE = "delete"

	MAX_LOOP_CHECK_OTP = 10

	OTP_FORGET_PASSWORD_TYPE      = 2
	OTP_REGISTER_SUB_ACCOUNT_TYPE = 3

	TYPE_BODY_MAIL = "text/html"

	LANGUAGE_ID_CODE = "id"
	LANGUAGE_EN_CODE = "en"

	FEATURE_TERMINAL_ID       = 11
	FEATURE_VERSION_ID        = 12
	FEATURE_BANK_PARTNER_ID   = 9
	FEATURE_BIN_ID            = 8
	FEATURE_ROLE_ID           = 7
	FEATURE_USER_ID           = 6
	FEATURE_MERCHANT_GROUP_ID = 2
	FEATURE_MERCHANT_ID       = 3

	ACTION_UPDATE = "UPDATE"
	ACTION_DELETE = "DELETE"
	ACTION_CREATE = "CREATE"
	ACTION_IMPORT = "IMPORT"

	HTTP_GET    = "GET"
	HTTP_POST   = "POST"
	HTTP_PUT    = "PUT"
	HTTP_DELETE = "DELETE"

	IMAGE_PNG_EXTENSION = "png"

	CONTENT_TYPE_IMAGE_PNG        = "image/png"
	CONTENT_TYPE_APPLICATION_JSON = "application/json"

	FORMAT_TIME_SYNC                 = "2006-01-02 15:04:05 -0700"
	FORMAT_TIME_SYNCWITHOUT_TIMEZONE = "2006-01-02 15:04:05"
	FORMAT_DATE                      = "20060102"
	FORMAT_TIME_MINUTE               = "2006-01-02 15:04"

	URL_SDH_API = "https://sandiegohills.co.id/SDHAPI/MobileApps_Services.svc"

	SKIP_VALIDATE_USER_REDIS = PREFIX_PROJECT + "_SKIP_VALIDATE_USER"

	DEFAULT_THIRD_PARTY_SCHEDULE_SYNC_REDIS = PREFIX_PROJECT + "_DEFAULT_THIRD_PARTY_SCHEDULE_SYNC"

	DEFAULT_QUERY_SOFT_DELETE                 = "deleted_at IS NULL"
	DEFAULT_SUBJECT_FORGET_PASSWORD_MAIL      = "Forget Password - San Diego Hills"
	DEFAULT_SUBJECT_REGISTER_SUB_ACCOUNT_MAIL = "Register Sub Account - San Diego Hills"
	DEFAULT_THIRD_PARTY_SCHEDULE_SYNC         = 15
	DEFAULT_NOTIFICATION_SCHEDULE             = 10
	DEFAULT_DOMAIN                            = "http://localhost:8080"
	DEFAULT_EXPIRED_PAYMENT                   = 300
	DEFAULT_KEY                               = "654321"

	ERROR_MESSAGE_00 = "Internal Server Error"
	ERROR_MESSAGE_01 = "User account is not registered"
	ERROR_MESSAGE_02 = "User account is not active"
	ERROR_MESSAGE_03 = "Username or password is wrong"
)
