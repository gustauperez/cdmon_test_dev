package hosting

type Hosting struct {
    ID        string   `json:"id"`
    Name      string   `json:"name"`
    Cores     string   `json:"cores"`
    Memory    string   `json:"memory"`
    Disc      string   `json:"disc"`
}
