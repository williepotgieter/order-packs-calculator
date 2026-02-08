package ports

// Server interface isused for driving adapters
type Server interface {
	Run() error
}
