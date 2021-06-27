### Parser
- Data
    - Column names list which will be used to generate key name for lookup
    - Key name of tag
- Methods
    - lookUpTag

### TagStore
- Data
    - File path which is used to be loaded for initilization
    - Column name which is for reading the tag value
- Methods
    - init
    - readTagByKey

### Processor
- Data
- Methods
    - read
    - process
    - write

### CSVFileOperator
- Data
- Methods
    - read
    - write