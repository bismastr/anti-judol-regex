package regex

type RegexService interface {
	GetRegexList() (*RegexResponse, error)
}

type RegexServiceImpl struct {
}

func NewRegexService() *RegexServiceImpl {
	return &RegexServiceImpl{}
}

type RegexResponse struct {
	TotalData int      `json:"totalData"`
	RegexList *[]Regex `json:"regexList"`
}

type Regex struct {
	Regex string `json:"regex"`
}

func (s *RegexServiceImpl) GetRegexList() (*RegexResponse, error) {
	response := &RegexResponse{
		TotalData: 10,
		RegexList: &[]Regex{
			{
				Regex: "m+\\s*[a4]+\\s*x+\\s*w+\\s*i+\\s*n+",
			},
			{
				Regex: "j+\\s*[a4]+\\s*c+\\s*k+\\s*p+\\s*[o0]+\\s*t+",
			},
			{
				Regex: "p+\\s*e+\\s*t+\\s*i+\\s*r+",
			},
			{
				Regex: "z+\\s*e+\\s*u+\\s*s+",
			},
			{
				Regex: "k+\\s*[a4]+\\s*k+\\s*e+\\s*k+",
			},
			{
				Regex: "g+\\s*[a4]+\\s*c+\\s*[o0]+\\s*r+",
			},
			{
				Regex: "g+\\s*u+\\s*a+\\s*c+\\s*[o0]+\\s*r+",
			},
			{
				Regex: "t+\\s*[e3]+\\s*r+\\s*p+\\s*[e3]+\\s*r+\\s*c+\\s*[a4]+\\s*y+\\s*[a4]+",
			},
			{
				Regex: "c+\\s*u+\\s*[a4]+\\s*n+",
			},
			{
				Regex: "",
			},
		},
	}

	return response, nil
}
