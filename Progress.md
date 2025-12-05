Open Redirect Tools TODO

capability :
- enable pipelining
- enable batch processing

redirx/
├── cmd/
│   └── redirx/
│       └── main.go
├── internal/
│   ├── runner/           # Orkestrator (Manage Worker Pool)
│   ├── options/          # Parsing Flag CLI (Input user)
│   ├── input/            # Handle input stream (File/Stdin)
│   └── output/           # Handle writing results (File/Console)
├── pkg/
│   ├── scanner/          # Core logic (HTTP Client & Detection)
│   ├── logger/           # UI Terminal (Colorful)
│   └── utils/            # Small helper functions
├── payloads/
│   └── payloads.txt      # List of payloads
├── go.mod
└── README.md


TODO :
- js/dom based open redirect
- hpp
- verify ssl