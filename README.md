# sap-api-integrations-service-agent-reads  
sap-api-integrations-service-agent-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API サービス代行データを取得するマイクロサービスです。  
sap-api-integrations-service-agent-reads には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-service-agent-reads は、オンプレミス版である（＝クラウド版ではない）SAPC4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/serviceagent/overview  

## 動作環境
sap-api-integrations-service-agent-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。   
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。   
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須） 

## クラウド環境での利用  
sap-api-integrations-service-agent-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-service-agent-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/serviceagent/overview
* APIサービス名(=baseURL): c4codataapi

## 本レポジトリ に 含まれる API名
sap-api-integrations-service-agent-reads には、次の API をコールするためのリソースが含まれています。  

* ServiceAgentCollection（サービスエージェント - サービスエージェント）


## API への 値入力条件 の 初期値
sap-api-integrations-service-agent-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.ServiceAgentCollection.ServiceAgentID（サービス代行ID）


## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"ServiceAgentCollection" が指定されています。    
  
```
	"api_schema": "ServiceAgent",
	"accepter": ["ServiceAgentCollection"],
	"service_agent_code": "7000001",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "ServiceAgent",
	"accepter": ["All"],
	"service_agent_code": "7000001",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetServiceAgent(serviceAgentID string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "ServiceAgentCollection":
			func() {
				c.ServiceAgentCollection(serviceAgentID)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP の サービス代行データ が取得された結果の JSON の例です。  
以下の項目のうち、"ObjectID" ～ "ETag" は、/SAP_API_Output_Formatter/type.go 内 の Type ServiceAgentCollection {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-service-agent-reads/SAP_API_Caller/caller.go#L53",
	"function": "sap-api-integrations-service-agent-reads/SAP_API_Caller.(*SAPAPICaller).ServiceAgentCollection",
	"level": "INFO",
	"message": [
		{
			"ObjectID": "00163E038C701ED298B0726C673837B9",
			"ServiceAgentID": "7000001",
			"ServiceAgentUUID": "00163E03-8C70-1ED2-98B0-726C673837B9",
			"LifeCycleStatusCode": "2",
			"LifeCycleStatusCodeText": "Active",
			"TitleCode": "",
			"TitleCodeText": "",
			"AcademicTitleCode": "",
			"AcademicTitleCodeText": "",
			"BusinessPartnerFormattedName": "Jason Chamberland",
			"FirstName": "Jason",
			"LastName": "Chamberland",
			"MiddleName": "",
			"NickName": "",
			"GenderCode": "0",
			"GenderCodeText": "Gender not known",
			"LanguageCode": "",
			"LanguageCodeText": "",
			"NationalityCountryCode": "",
			"NationalityCountryCodeText": "",
			"BirthName": "",
			"FormattedPostalAddressDescription": "256 West El Camino Real / Mountian View 94041 / US",
			"CountryCode": "US",
			"CountryCodeText": "United States",
			"RegionCode": "CA",
			"RegionCodeText": "California",
			"AddressLine1": "",
			"AddressLine2": "",
			"HouseNumber": "256",
			"Street": "West El Camino Real",
			"AddressLine4": "",
			"AddressLine5": "",
			"District": "",
			"City": "Mountian View",
			"DifferentCity": "",
			"PostalCode": "94041",
			"County": "",
			"CompanyPostalCode": "",
			"POBox": "",
			"POBoxPostalCode": "",
			"POBoxCountryCode": "",
			"POBoxCountryCodeText": "",
			"POBoxRegionCode": "",
			"POBoxRegionCodeText": "",
			"POBoxCity": "",
			"TimeZoneCode": "PST",
			"TimeZoneCodeText": "(UTC-08:00) Pacific Time (Los Angeles, Tijuana, Yukon)",
			"TaxJurisdictionCode": "",
			"TaxJurisdictionCodeText": "",
			"Building": "",
			"Floor": "",
			"Room": "",
			"InhouseMail": "",
			"OfficePhoneNumber": "+1 6503275561",
			"MobilePhoneNumber": "+1 4086275100",
			"FaxNumber": "+1 4082953819",
			"Email": "Jason.Chamberland@ondemand.com",
			"NormalisedOfficePhoneNumber": "+16503275561",
			"NormalisedMobilePhoneNumber": "+14086275100",
			"CreatedOn": "2013-01-19T00:28:39+09:00",
			"CreatedBy": "Eddie Smoke",
			"CreatedByIdentityUUID": "00163E03-A070-1EE2-88BA-39BD20F290B5",
			"ChangedOn": "2016-07-13T20:41:32+09:00",
			"ChangedBy": "Technical User SAP_SYSTEM",
			"ChangedByIdentityUUID": "00163E03-A06A-1EE2-87C4-A680EDD73F8D",
			"EntityLastChangedOn": "2016-07-13T20:41:32+09:00",
			"ETag": "2017-08-10T23:01:59+09:00"
		}
	],
	"time": "2022-08-10T12:48:35+09:00"
}

```