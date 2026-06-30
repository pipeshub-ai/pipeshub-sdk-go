# RecordTypeEnum

Type of content. Mirrors the backend `RecordType` enum
(`backend/python/app/models/entities.py`); connector-sourced records
may use any of the connector-specific types below.
- FILE: Uploaded or synced documents (PDF, DOCX, etc.)
- DRIVE: Drive/folder container (Google Drive, OneDrive, etc.)
- WEBPAGE: Web pages crawled or bookmarked
- DATABASE: Database object (e.g. Notion database)
- DATASOURCE: Data source object
- MESSAGE: Chat/messaging content (Slack, Teams)
- MAIL: Email messages (Gmail, Outlook)
- GROUP_MAIL: Group/shared mailbox email messages
- TICKET: Support/issue tickets (Jira, ServiceNow)
- COMMENT: Comments from collaboration tools
- INLINE_COMMENT: Inline comments anchored to content (e.g. Confluence)
- CONFLUENCE_PAGE: Confluence page
- CONFLUENCE_BLOGPOST: Confluence blog post
- SHAREPOINT_PAGE: SharePoint page
- SHAREPOINT_LIST: SharePoint list
- SHAREPOINT_LIST_ITEM: SharePoint list item
- SHAREPOINT_DOCUMENT_LIBRARY: SharePoint document library
- LINK: Web link / bookmark
- PROJECT: Project entity (e.g. Jira project)
- PULL_REQUEST: Source-control pull request
- MEETING: Meeting record (e.g. Zoom)
- PRODUCT: Product entity (CRM)
- DEAL: Deal/opportunity entity (CRM)
- CASE: Case entity (CRM/support)
- TASK: Task entity
- ARTIFACT: Generated/derived artifact
- CODE_FILE: Source-code file
- SQL_TABLE: SQL table object
- SQL_VIEW: SQL view object
- OTHERS: Miscellaneous content types



## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `RecordTypeEnumFile`                      | FILE                                      |
| `RecordTypeEnumDrive`                     | DRIVE                                     |
| `RecordTypeEnumWebpage`                   | WEBPAGE                                   |
| `RecordTypeEnumDatabase`                  | DATABASE                                  |
| `RecordTypeEnumDatasource`                | DATASOURCE                                |
| `RecordTypeEnumMessage`                   | MESSAGE                                   |
| `RecordTypeEnumMail`                      | MAIL                                      |
| `RecordTypeEnumGroupMail`                 | GROUP_MAIL                                |
| `RecordTypeEnumTicket`                    | TICKET                                    |
| `RecordTypeEnumComment`                   | COMMENT                                   |
| `RecordTypeEnumInlineComment`             | INLINE_COMMENT                            |
| `RecordTypeEnumConfluencePage`            | CONFLUENCE_PAGE                           |
| `RecordTypeEnumConfluenceBlogpost`        | CONFLUENCE_BLOGPOST                       |
| `RecordTypeEnumSharepointPage`            | SHAREPOINT_PAGE                           |
| `RecordTypeEnumSharepointList`            | SHAREPOINT_LIST                           |
| `RecordTypeEnumSharepointListItem`        | SHAREPOINT_LIST_ITEM                      |
| `RecordTypeEnumSharepointDocumentLibrary` | SHAREPOINT_DOCUMENT_LIBRARY               |
| `RecordTypeEnumLink`                      | LINK                                      |
| `RecordTypeEnumProject`                   | PROJECT                                   |
| `RecordTypeEnumPullRequest`               | PULL_REQUEST                              |
| `RecordTypeEnumMeeting`                   | MEETING                                   |
| `RecordTypeEnumProduct`                   | PRODUCT                                   |
| `RecordTypeEnumDeal`                      | DEAL                                      |
| `RecordTypeEnumCase`                      | CASE                                      |
| `RecordTypeEnumTask`                      | TASK                                      |
| `RecordTypeEnumArtifact`                  | ARTIFACT                                  |
| `RecordTypeEnumCodeFile`                  | CODE_FILE                                 |
| `RecordTypeEnumSQLTable`                  | SQL_TABLE                                 |
| `RecordTypeEnumSQLView`                   | SQL_VIEW                                  |
| `RecordTypeEnumOthers`                    | OTHERS                                    |