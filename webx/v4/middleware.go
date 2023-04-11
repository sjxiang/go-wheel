package webx

type Middleware func(next HandleFunc) HandleFunc