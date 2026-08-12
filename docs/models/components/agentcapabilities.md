# AgentCapabilities

Per-request agent capability toggles. Only meaningful when `chatMode`
selects an agent mode; ignored otherwise. Each field falls back to its
own `default` below when omitted — a missing flag is not uniformly
`true`. Omitting the whole object applies every default.



## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `InternalSearch`                                                          | **bool*                                                                   | :heavy_minus_sign:                                                        | Whether the agent may search internal knowledge bases for this turn.      |
| `WebSearch`                                                               | **bool*                                                                   | :heavy_minus_sign:                                                        | Whether the agent may perform web search for this turn.                   |
| `DeepSearch`                                                              | **bool*                                                                   | :heavy_minus_sign:                                                        | Whether the agent may use deeper, higher-latency retrieval for this turn. |