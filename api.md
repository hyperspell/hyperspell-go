# Shared Response Types

- <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go/shared#Metadata">Metadata</a>
- <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go/shared#Notification">Notification</a>
- <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go/shared#QueryResult">QueryResult</a>
- <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go/shared#Resource">Resource</a>

# Connections

Response Types:

- <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#ConnectionListResponse">ConnectionListResponse</a>
- <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#ConnectionRevokeResponse">ConnectionRevokeResponse</a>

Methods:

- <code title="get /connections/list">client.Connections.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#ConnectionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#ConnectionListResponse">ConnectionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /connections/{connection_id}/revoke">client.Connections.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#ConnectionService.Revoke">Revoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, connectionID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#ConnectionRevokeResponse">ConnectionRevokeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Folders

Response Types:

- <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#FolderListResponse">FolderListResponse</a>
- <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#FolderDeletePolicyResponse">FolderDeletePolicyResponse</a>
- <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#FolderListPoliciesResponse">FolderListPoliciesResponse</a>
- <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#FolderSetPoliciesResponse">FolderSetPoliciesResponse</a>

Methods:

- <code title="get /connections/{connection_id}/folders">client.Folders.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#FolderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, connectionID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#FolderListParams">FolderListParams</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#FolderListResponse">FolderListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /connections/{connection_id}/folder-policies/{policy_id}">client.Folders.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#FolderService.DeletePolicy">DeletePolicy</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, policyID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#FolderDeletePolicyParams">FolderDeletePolicyParams</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#FolderDeletePolicyResponse">FolderDeletePolicyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /connections/{connection_id}/folder-policies">client.Folders.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#FolderService.ListPolicies">ListPolicies</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, connectionID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#FolderListPoliciesResponse">FolderListPoliciesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /connections/{connection_id}/folder-policies">client.Folders.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#FolderService.SetPolicies">SetPolicies</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, connectionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#FolderSetPoliciesParams">FolderSetPoliciesParams</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#FolderSetPoliciesResponse">FolderSetPoliciesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Integrations

Response Types:

- <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#IntegrationListResponse">IntegrationListResponse</a>
- <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#IntegrationConnectResponse">IntegrationConnectResponse</a>

Methods:

- <code title="get /integrations/list">client.Integrations.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#IntegrationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#IntegrationListResponse">IntegrationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /integrations/{integration_id}/connect">client.Integrations.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#IntegrationService.Connect">Connect</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, integrationID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#IntegrationConnectParams">IntegrationConnectParams</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#IntegrationConnectResponse">IntegrationConnectResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## GoogleCalendar

Response Types:

- <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#Calendar">Calendar</a>

Methods:

- <code title="get /integrations/google_calendar/list">client.Integrations.GoogleCalendar.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#IntegrationGoogleCalendarService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#Calendar">Calendar</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## WebCrawler

Response Types:

- <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#IntegrationWebCrawlerIndexResponse">IntegrationWebCrawlerIndexResponse</a>

Methods:

- <code title="get /integrations/web_crawler/index">client.Integrations.WebCrawler.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#IntegrationWebCrawlerService.Index">Index</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#IntegrationWebCrawlerIndexParams">IntegrationWebCrawlerIndexParams</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#IntegrationWebCrawlerIndexResponse">IntegrationWebCrawlerIndexResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Slack

Response Types:

- <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#IntegrationSlackListResponse">IntegrationSlackListResponse</a>

Methods:

- <code title="get /integrations/slack/list">client.Integrations.Slack.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#IntegrationSlackService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#IntegrationSlackListParams">IntegrationSlackListParams</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#IntegrationSlackListResponse">IntegrationSlackListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Memories

Response Types:

- <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#Memory">Memory</a>
- <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#MemoryStatus">MemoryStatus</a>
- <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#MemoryDeleteResponse">MemoryDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#MemoryAddBulkResponse">MemoryAddBulkResponse</a>
- <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#MemoryStatusResponse">MemoryStatusResponse</a>

Methods:

- <code title="post /memories/update/{source}/{resource_id}">client.Memories.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#MemoryService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, resourceID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#MemoryUpdateParams">MemoryUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#MemoryStatus">MemoryStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /memories/list">client.Memories.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#MemoryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#MemoryListParams">MemoryListParams</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go/packages/pagination#CursorPage">CursorPage</a>[<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go/shared#Resource">Resource</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /memories/delete/{source}/{resource_id}">client.Memories.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#MemoryService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, resourceID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#MemoryDeleteParams">MemoryDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#MemoryDeleteResponse">MemoryDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /memories/add">client.Memories.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#MemoryService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#MemoryAddParams">MemoryAddParams</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#MemoryStatus">MemoryStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /memories/add/bulk">client.Memories.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#MemoryService.AddBulk">AddBulk</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#MemoryAddBulkParams">MemoryAddBulkParams</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#MemoryAddBulkResponse">MemoryAddBulkResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /memories/get/{source}/{resource_id}">client.Memories.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#MemoryService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, resourceID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#MemoryGetParams">MemoryGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#Memory">Memory</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /memories/query">client.Memories.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#MemoryService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#MemorySearchParams">MemorySearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go/shared#QueryResult">QueryResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /memories/status">client.Memories.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#MemoryService.Status">Status</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#MemoryStatusResponse">MemoryStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /memories/upload">client.Memories.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#MemoryService.Upload">Upload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#MemoryUploadParams">MemoryUploadParams</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#MemoryStatus">MemoryStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Evaluate

Response Types:

- <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#EvaluateListQueriesResponse">EvaluateListQueriesResponse</a>
- <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#EvaluateScoreHighlightResponse">EvaluateScoreHighlightResponse</a>
- <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#EvaluateScoreQueryResponse">EvaluateScoreQueryResponse</a>

Methods:

- <code title="get /evaluate/query/{query_id}">client.Evaluate.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#EvaluateService.GetQuery">GetQuery</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, queryID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go/shared#QueryResult">QueryResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /evaluate/queries">client.Evaluate.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#EvaluateService.ListQueries">ListQueries</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#EvaluateListQueriesParams">EvaluateListQueriesParams</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go/packages/pagination#CursorPage">CursorPage</a>[<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#EvaluateListQueriesResponse">EvaluateListQueriesResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /evaluate/highlight/{highlight_id}">client.Evaluate.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#EvaluateService.ScoreHighlight">ScoreHighlight</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, highlightID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#EvaluateScoreHighlightParams">EvaluateScoreHighlightParams</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#EvaluateScoreHighlightResponse">EvaluateScoreHighlightResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /evaluate/query/{query_id}">client.Evaluate.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#EvaluateService.ScoreQuery">ScoreQuery</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, queryID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#EvaluateScoreQueryParams">EvaluateScoreQueryParams</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#EvaluateScoreQueryResponse">EvaluateScoreQueryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Actions

Response Types:

- <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#ActionAddReactionResponse">ActionAddReactionResponse</a>
- <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#ActionSendMessageResponse">ActionSendMessageResponse</a>

Methods:

- <code title="post /actions/add_reaction">client.Actions.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#ActionService.AddReaction">AddReaction</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#ActionAddReactionParams">ActionAddReactionParams</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#ActionAddReactionResponse">ActionAddReactionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /actions/send_message">client.Actions.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#ActionService.SendMessage">SendMessage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#ActionSendMessageParams">ActionSendMessageParams</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#ActionSendMessageResponse">ActionSendMessageResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Sessions

Methods:

- <code title="post /trace/add">client.Sessions.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#SessionService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#SessionAddParams">SessionAddParams</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#MemoryStatus">MemoryStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Vaults

Response Types:

- <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#VaultListResponse">VaultListResponse</a>

Methods:

- <code title="get /vault/list">client.Vaults.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#VaultService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#VaultListParams">VaultListParams</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go/packages/pagination#CursorPage">CursorPage</a>[<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#VaultListResponse">VaultListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Auth

Response Types:

- <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#Token">Token</a>
- <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#AuthDeleteUserResponse">AuthDeleteUserResponse</a>
- <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#AuthMeResponse">AuthMeResponse</a>

Methods:

- <code title="delete /auth/delete">client.Auth.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#AuthService.DeleteUser">DeleteUser</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#AuthDeleteUserResponse">AuthDeleteUserResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /auth/me">client.Auth.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#AuthService.Me">Me</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#AuthMeResponse">AuthMeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /auth/user_token">client.Auth.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#AuthService.UserToken">UserToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#AuthUserTokenParams">AuthUserTokenParams</a>) (\*<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go">hyperspell</a>.<a href="https://pkg.go.dev/github.com/hyperspell/hyperspell-go#Token">Token</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
