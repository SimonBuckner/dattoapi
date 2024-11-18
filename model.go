package dattoapi

import "time"

// Types taken from https://portal.dattobackup.com/integrations/api

type Pagination struct {
	Page       int64 `json:"page"`
	PerPage    int64 `json:"perPage"`
	TotalPages int64 `json:"totalPages"`
	Count      int64 `json:"count"`
}

type Device struct {
	SerialNumber          string  `json:"serialNumber"`
	Name                  string  `json:"name"`
	Region                *Region `json:"region"`
	Model                 string  `json:"model"`
	LastSeenDate          *string `json:"lastSeenDate"`
	OrganizationId        *int64  `json:"organizationId"`
	ClientCompanyName     *string `json:"clientCompanyName"`
	OrganizationName      *string `json:"organizationName"`
	Hidden                *bool   `json:"hidden"`
	ServicePlan           *string `json:"servicePlan"`
	RegistrationDate      string  `json:"registrationDate"`
	ServicePeriod         string  `json:"servicePeriod"`
	WarrantyExpire        string  `json:"warrantyExpire"`
	LocalStorageUsed      Storage `json:"localStorageUsed"`
	LocalStorageAvailable Storage `json:"localStorageAvailable"`
	OffsiteStorageUsed    Storage `json:"offsiteStorageUsed"`
	TotalManagedDisk      Storage `json:"totalManagedDisk"`
	InternalIP            string  `json:"internalIP"`
	ActiveTickets         int64   `json:"activeTickets"`
	ResellerCompanyName   string  `json:"resellerCompanyName"`
	AgentCount            int64   `json:"agentCount"`
	ShareCount            int64   `json:"shareCount"`
	AlertCount            int64   `json:"alertCount"`
	Uptime                int64   `json:"uptime"`
	RemoteWebUrl          string  `json:"remoteWebUrl"`
}

type ProtectedMachine struct {
	Mac    *string `json:"mac"`
	Serial *string `json:"serial"`
}

type Storage struct {
	Size  int64  `json:"size"`
	Units string `json:"units"`
}

type Region struct {
	Location string  `json:"location"`
	Provider *string `json:"provider"`
}

type DeviceAsset struct {
	Name                        string           `json:"name"`
	AssetId                     int64            `json:"assetId"`
	Volume                      *string          `json:"volume"`
	LocalIp                     *string          `json:"localIp"`
	Os                          *string          `json:"os"`
	ProtectedVolumesCount       int64            `json:"protectedVolumesCount"`
	UnprotectedVolumesCount     int64            `json:"unprotectedVolumesCount"`
	ProtectedVolumeNames        []string         `json:"protectedVolumeNames"`
	UnprotectedVolumeNames      []string         `json:"unprotectedVolumeNames"`
	AgentVersion                string           `json:"agentVersion"`
	IsPaused                    bool             `json:"isPaused"`
	IsArchived                  bool             `json:"isArchived"`
	LatestOffsite               *int64           `json:"latestOffsite"`
	LocalSnapshots              int64            `json:"localSnapshots"`
	LastSnapshot                int64            `json:"lastSnapshot"`
	LastScreenshotAttempt       int64            `json:"lastScreenshotAttempt"`
	LastScreenshotAttemptStatus bool             `json:"lastScreenshotAttemptStatus"`
	LastScreenshotUrl           string           `json:"lastScreenshotUrl"`
	Fqdn                        *string          `json:"fqdn"`
	Backups                     []Snapshot       `json:"backups"`
	ProtectedMachine            ProtectedMachine `json:"protectedMachine"`
}

type Snapshot struct {
	Timestamp            time.Time            `json:"timestamp*"`
	Backup               Backup               `json:"backup*"`
	LocalVerification    LocalVerification    `json:"localVerification"`
	AdvancedVerification AdvancedVerification `json:"advancedVerification"`
}

type Backup struct {
	Status       **string `json:"status*"`
	ErrorMessage **string `json:"errorMessage*"`
}

type LocalVerification struct {
	Status *string                   `json:"status*"`
	Errors *[]LocalVerificationError `json:"errors*"`
}

type AdvancedVerification struct {
	ScreenshotVerification ScreenshotVerification `json:"screenshotVerification"`
}

type ScreenshotVerification struct {
	Status string `json:"status"`
	Image  string `json:"image"`
}

type LocalVerificationError struct {
	ErrorType    *string `json:"errorType*"`
	ErrorMessage *string `json:"errorMessage*"`
}

type DeviceVMRestore struct {
	RestoreId   int64   `json:"restoreId"`
	AssetId     *int64  `json:"assetId"`
	CpuCount    int64   `json:"cpuCount"`
	RamSize     int64   `json:"ramSize"`
	SocketCount int64   `json:"socketCount"`
	Timezone    *string `json:"timezone"`
	NetworkName *string `json:"networkName"`
	StorageBus  string  `json:"storageBus"`
	ModelType   string  `json:"modelType"`
	MacAddress  string  `json:"macAddress"`
}

type Alert struct {
	Type          string `json:"type"`
	Threshold     int64  `json:"threshold"`
	Unit          string `json:"unit"`
	DateTriggered string `json:"dateTriggered"`
	DateSent      string `json:"dateSent"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ActivityLog struct {
	Timestamp         string   `json:"timestamp"`
	RequestId         string   `json:"requestId"`
	TargetType        string   `json:"targetType"`
	TargetId          string   `json:"targetId"`
	TargetDisplayName string   `json:"targetDisplayName"`
	ClientName        *string  `json:"clientName"`
	Interface         string   `json:"interface"`
	User              string   `json:"user"`
	UserRoles         []string `json:"userRoles"`
	IpAddress         string   `json:"ipAddress"`
	Action            string   `json:"action"`
	MessageEN         string   `json:"messageEN"`
	Success           bool     `json:"success"`
}

type BcdrAgent struct {
	Uuid              string           `json:"uuid"`
	ShortCode         string           `json:"shortCode"`
	Type              string           `json:"type"`
	Hostname          string           `json:"hostname"`
	LastSnapshot      int64            `json:"lastSnapshot"`
	LastScreenshot    int64            `json:"lastScreenshot"`
	ScreenshotSuccess bool             `json:"screenshotSuccess"`
	ProtectedMachine  ProtectedMachine `json:"protectedMachine"`
}

type BcdrClient struct {
	OrganizationId   *int64      `json:"organizationId"`
	OrganizationName string      `json:"organizationName"`
	ClientName       string      `json:"clientName"`
	Agents           []BcdrAgent `json:"agents"`
}

type SaasProtection struct {
	CustomerName           string          `json:"customerName"`
	ProtectedDomainName    string          `json:"protectedDomainName"`
	ProductType            string          `json:"productType"`
	ExternalSubscriptionId string          `json:"externalSubscriptionId"`
	RetentionType          string          `json:"retentionType"`
	BackupSuccessRate      int64           `json:"backupSuccessRate"`
	TotalProtectedBytes    int64           `json:"totalProtectedBytes"`
	Services               []SaasService   `json:"services"`
	SeatUsages             []SaasSeatUsage `json:"seatUsages"`
	Seats                  []SaasSeat      `json:"seats"`
}
type SaasSeatUsage struct {
	SeatType    string `json:"seatType"`
	Active      int64  `json:"active"`
	Paused      int64  `json:"paused"`
	Archived    int64  `json:"archived"`
	Unprotected int64  `json:"unprotected"`
	AutoAdd     string `json:"autoAdd"`
}

type SaasService struct {
	AppType                string              `json:"appType"`
	BackupHistory          []BackupHistoryItem `json:"backupHistory"`
	LastFullyProtectedTime string              `json:"lastFullyProtectedTime"`
	UningestedServiceCount int64               `json:"uningestedServiceCount"`
	UsedBytes              int64               `json:"usedBytes"`
}

type SaasSeat struct {
	MainId    string `json:"mainId"`
	Name      string `json:"name"`
	SeatType  string `json:"seatType"`
	SeatState string `json:"seatState"`
	Billable  string `json:"billable"`
	DateAdded string `json:"dateAdded"`
}

type BackupStats struct {
	ActiveServicesCount                 int64 `json:"activeServicesCount"`
	ActiveServicesWithRecentBackupCount int64 `json:"activeServicesWithRecentBackupCount"`
	BackupPercentage                    int64 `json:"backupPercentage"`
}
type SaasDomain struct {
	BackupStats            BackupStats `json:"backupStats"`
	Domain                 string      `json:"domain"`
	SaasCustomerId         int64       `json:"saasCustomerId"`
	SaasCustomerName       string      `json:"saasCustomerName"`
	OrganizationId         *int64      `json:"organizationId"`
	OrganizationName       *string     `json:"organizationName"`
	SeatsUsed              int64       `json:"seatsUsed"`
	ProductType            string      `json:"productType"`
	ExternalSubscriptionId string      `json:"externalSubscriptionId"`
	RetentionType          string      `json:"retentionType"`
}

type BackupHistoryItem struct {
	ActiveServiceCount                  int64  `json:"activeServiceCount"`
	ActiveServiceWithBackupCount        int64  `json:"activeServiceWithBackupCount"`
	ActiveServiceWithPerfectBackupCount int64  `json:"activeServiceWithPerfectBackupCount"`
	EndTime                             int64  `json:"endTime"`
	StartTime                           int64  `json:"startTime"`
	Status                              string `json:"status"`
	TimeWindow                          string `json:"timeWindow"`
	TotalServiceCount                   int64  `json:"totalServiceCount"`
}
type BulkSeatAssignment struct {
	Action     string `json:"action"`
	AppType    string `json:"appType"`
	CustomerId int64  `json:"customerId"`
	Id         int64  `json:"id"`
	Status     string `json:"status"`
}

type AppTypeItem struct {
	AppType       string              `json:"appType"`
	BackupHistory []BackupHistoryItem `json:"backupHistory"`
}

type BackupReportSuiteItem struct {
	SuiteType string        `json:"suiteType"`
	AppTypes  []AppTypeItem `json:"appTypes"`
}

type BackupReportRemoteIdItem struct {
	Active      int64    `json:"active"`
	Archived    int64    `json:"archived"`
	Discovered  int64    `json:"discovered"`
	Paused      int64    `json:"paused"`
	Unprotected int64    `json:"unprotected"`
	SeatType    string   `json:"seatType"`
	RemoteIds   []string `json:"remoteIds"`
}
type BackupReportItem struct {
	CustomerId   int64                      `json:"customerId"`
	CustomerName string                     `json:"customerName"`
	Suites       []BackupReportSuiteItem    `json:"suites"`
	UsedBytes    int64                      `json:"usedBytes"`
	RemoteIds    []BackupReportRemoteIdItem `json:"remoteIds"`
}
type SaasBackup struct {
	Pagination Pagination         `json:"pagination"`
	Items      []BackupReportItem `json:"items"`
}
type CustomerBackupStats struct {
	BackupStatusLastSevenDays      BackupStatusLastSevenDays `json:"backupStatusLastSevenDays"`
	BackupStatusLastSevenDaysTotal []string                  `json:"backupStatusLastSevenDaysTotal"`
	LastBackups                    LastBackups               `json:"lastBackups"`
	NumberOfStandardLicensesTotal  int64                     `json:"numberOfStandardLicensesTotal"`
	NumberOfUsers                  int64                     `json:"numberOfUsers"`
	Region                         string                    `json:"region"`
	StorageInformation             StorageInformation        `json:"storageInformation"`
	TenantId                       string                    `json:"tenantId"`
	TenantName                     string                    `json:"tenantName"`
	Users                          []BackupStatsUserItem     `json:"users"`
}
type BackupStatusLastSevenDays struct {
	TotalForCalendar   string `json:"totalForCalendar"`
	TotalForContact    string `json:"totalForContact"`
	TotalForDrive      string `json:"totalForDrive"`
	TotalForMail       string `json:"totalForMail"`
	TotalForSharePoint string `json:"totalForSharePoint"`
	TotalForSites      string `json:"totalForSites"`
	TotalForTeams      string `json:"totalForTeams"`
}
type UserBackupStatusLastSevenDays struct {
	TotalForCalendar string `json:"totalForCalendar"`
	TotalForContact  string `json:"totalForContact"`
	TotalForDrive    string `json:"totalForDrive"`
	TotalForMail     string `json:"totalForMail"`
}

type LastBackups struct {
	Calendar      []CustomerBackupStatsItem `json:"calendar"`
	Contact       []CustomerBackupStatsItem `json:"contact"`
	DomainBackups []CustomerBackupStatsItem `json:"domainBackups"`
	Drive         []CustomerBackupStatsItem `json:"drive"`
	Mail          []CustomerBackupStatsItem `json:"mail"`
	SharePoint    []CustomerBackupStatsItem `json:"sharePoint"`
	Sites         []CustomerBackupStatsItem `json:"sites"`
	Teams         []CustomerBackupStatsItem `json:"teams"`
}
type UserLastBackups struct {
	LastForCalendar string `json:"lastForCalendar"`
	LastForContact  string `json:"lastForContact"`
	LastForDrive    string `json:"lastForDrive"`
	LastForMail     string `json:"lastForMail"`
}

type CustomerBackupStatsItem struct {
	ProblemsCount int64  `json:"problemsCount"`
	Status        string `json:"status"`
	Timestamp     int64  `json:"timestamp"`
}

type BackupStatsUserItem struct {
	BackupStatusLastSevenDays      UserBackupStatusLastSevenDays `json:"backupStatusLastSevenDays"`
	Email                          string                        `json:"email"`
	Id                             string                        `json:"id"`
	LastBackupStatus               UserLastBackups               `json:"lastBackupStatus"`
	LastBackupStatusSevenDaysTotal string                        `json:"lastBackupStatusSevenDaysTotal"`
	LastBackupStatusTotal          string                        `json:"lastBackupStatusTotal"`
	LastBackupTimestamp            LastBackupTimestamp           `json:"lastBackupTimestamp"`
	LastBackupTimestampTotal       int64                         `json:"lastBackupTimestampTotal"`
	LastBackupsStatusTotal         []CustomerBackupStatsItem     `json:"lastBackupsStatusTotal"`
	StorageInformation             StorageInformation            `json:"storageInformation"`
}

type LastBackupTimestamp struct {
	LastCalendarBackupTimestamp int64 `json:"lastCalendarBackupTimestamp"`
	LastContactBackupTimestamp  int64 `json:"lastContactBackupTimestamp"`
	LastDriveBackupTimestamp    int64 `json:"lastDriveBackupTimestamp"`
	LastMailBackupTimestamp     int64 `json:"lastMailBackupTimestamp"`
}

type Last10Backups struct {
	Timestamp         int64 `json:"timestamp"`
	BackupStatus      bool  `json:"backupStatus"`
	ScreenshotSuccess bool  `json:"screenshotSuccess"`
}

type DTCAssetsData struct {
	Name                  string          `json:"name"`
	Status                string          `json:"status"`
	Type                  string          `json:"type"`
	LastBackup            int64           `json:"lastBackup"`
	ScreenshotSuccess     bool            `json:"screenshotSuccess"`
	AssetUuid             string          `json:"assetUuid"`
	OrganizationId        int64           `json:"organizationId"`
	OrganizationName      string          `json:"organizationName"`
	ResellerId            int64           `json:"resellerId"`
	BackupStateSuccessful bool            `json:"backupStateSuccessful"`
	LastScreenshotUrl     string          `json:"lastScreenshotUrl"`
	Last10Backups         []Last10Backups `json:"last10Backups"`
	Hostname              string          `json:"hostname"`
	MacAddress            string          `json:"macAddress"`
	ExternalIp            string          `json:"externalIp"`
	SmbiosVendor          string          `json:"smbiosVendor"`
	SmbiosName            string          `json:"smbiosName"`
	AgentCode             string          `json:"agentCode"`
	ProtectedVolume       int64           `json:"protectedVolume"`
}

type DTCAssetDetails struct {
	Name            string          `json:"name"`
	AssetUuid       string          `json:"assetUuid"`
	LastCheckin     int64           `json:"lastCheckin"`
	LastBackup      int64           `json:"lastBackup"`
	LastScreenshot  int64           `json:"lastScreenshot"`
	Subscription    string          `json:"subscription"`
	AgentCode       string          `json:"agentCode"`
	ProtectedVolume int64           `json:"protectedVolume"`
	Last10Backups   []Last10Backups `json:"last10Backups"`
	Hostname        string          `json:"hostname"`
	MacAddress      string          `json:"macAddress"`
	ExternalIp      string          `json:"externalIp"`
	SmbiosVendor    string          `json:"smbiosVendor"`
	SmbiosName      string          `json:"smbiosName"`
}
type RMMTemplates struct {
	Type                    string `json:"type"`
	TokenUuid               string `json:"tokenUuid"`
	ResellerId              int64  `json:"resellerId"`
	ClientId                int64  `json:"clientId"`
	RegionId                int64  `json:"regionId"`
	SubscriptionUuid        string `json:"subscriptionUuid"`
	CreateDate              string `json:"createDate"`
	ExpireDate              string `json:"expireDate"`
	LastUsed                string `json:"lastUsed"`
	TimesUsed               int64  `json:"timesUsed"`
	CumulativeTimesUsed     int64  `json:"cumulativeTimesUsed"`
	LastRotationDate        string `json:"lastRotationDate"`
	Active                  bool   `json:"active"`
	EssentialsSubsciptionId int64  `json:"essentialsSubsciptionId"`
}

type StorageInformation struct {
	ProtectedBytes         int64 `json:"protectedBytes"`
	ProtectedCalendarBytes int64 `json:"protectedCalendarBytes"`
	ProtectedContactBytes  int64 `json:"protectedContactBytes"`
	ProtectedDriveBytes    int64 `json:"protectedDriveBytes"`
	ProtectedMailBytes     int64 `json:"protectedMailBytes"`
	UsedBytes              int64 `json:"usedBytes"`
	UsedCalendarBytes      int64 `json:"usedCalendarBytes"`
	UsedContactBytes       int64 `json:"usedContactBytes"`
	UsedDriveBytes         int64 `json:"usedDriveBytes"`
	UsedMailBytes          int64 `json:"usedMailBytes"`
}
type BandwidthSettingsRequest struct {
	PauseWhileMetered               bool   `json:"pauseWhileMetered"`
	MaximumBandwidthInBits          *int64 `json:"maximumBandwidthInBits"`
	MaximumThrottledBandwidthInBits *int64 `json:"maximumThrottledBandwidthInBits"`
}
type BandwidthSettingsResponse struct {
	PauseWhileMetered               bool   `json:"pauseWhileMetered"`
	MaximumBandwidthInBits          *int64 `json:"maximumBandwidthInBits"`
	MaximumThrottledBandwidthInBits *int64 `json:"maximumThrottledBandwidthInBits"`
	AgentUuid                       string `json:"agentUuid"`
}
