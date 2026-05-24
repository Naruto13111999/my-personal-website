package data

type TechBadge struct {
	Name  string
	Icon  string
	Color string
}

var techMeta = map[string]TechBadge{
	"Golang":           {Name: "Golang", Icon: iconLocal("golang"), Color: "#00ADD8"},
	"Go":               {Name: "Go", Icon: iconLocal("golang"), Color: "#00ADD8"},
	"C++":              {Name: "C++", Icon: iconLocal("cplusplus"), Color: "#00599C"},
	"C/C++":            {Name: "C/C++", Icon: iconLocal("cplusplus"), Color: "#00599C"},
	"C":                {Name: "C", Icon: iconDevicon("c", "original"), Color: "#A8B9CC"},
	"Python":           {Name: "Python", Icon: iconDevicon("python", "original"), Color: "#3776AB"},
	"PostgreSQL":       {Name: "PostgreSQL", Icon: iconLocal("postgresql"), Color: "#4169E1"},
	"Redis":            {Name: "Redis", Icon: iconDevicon("redis", "original"), Color: "#DC382D"},
	"gRPC":             {Name: "gRPC", Icon: iconLocal("grpc"), Color: "#244C5A"},
	"Docker":           {Name: "Docker", Icon: iconLocal("docker"), Color: "#2496ED"},
	"Kubernetes":       {Name: "Kubernetes", Icon: iconLocal("kubernetes"), Color: "#326CE5"},
	"AWS":              {Name: "AWS", Icon: iconLocal("aws"), Color: "#FF9900"},
	"Amazon SQS":       {Name: "Amazon SQS", Icon: iconLocal("aws"), Color: "#FF9900"},
	"Protocol Buffers": {Name: "Protocol Buffers", Icon: iconDevicon("google", "original"), Color: "#4285F4"},
	"REST API":         {Name: "REST API", Icon: iconDevicon("nginx", "original"), Color: "#009639"},
	"CI/CD":            {Name: "CI/CD", Icon: iconDevicon("githubactions", "plain"), Color: "#2088FF"},
	"NFS":              {Name: "NFS", Icon: iconLocal("nfs"), Color: "#FCC624"},
	"TCP/IP":           {Name: "TCP/IP", Icon: iconDevicon("nginx", "original"), Color: "#009639"},
	"Object Storage":   {Name: "Object Storage", Icon: iconLocal("aws"), Color: "#FF9900"},
	"Solace":           {Name: "Solace", Icon: "", Color: "#FF6B35"},
	"HFT":              {Name: "HFT", Icon: "", Color: "#F59E0B"},
	"Low Latency":      {Name: "Low Latency", Icon: "", Color: "#EF4444"},
	"Market Feed":      {Name: "Market Feed", Icon: "", Color: "#8B5CF6"},
	"Windows MDM":      {Name: "Windows MDM", Icon: iconDevicon("windows8", "original"), Color: "#0078D4"},
	"Windows Policy":   {Name: "Windows Policy", Icon: iconDevicon("windows8", "original"), Color: "#0078D4"},
	"Android EMM":      {Name: "Android EMM", Icon: iconDevicon("android", "original"), Color: "#3DDC84"},
	"SCEP":             {Name: "SCEP", Icon: iconDevicon("openssl", "plain"), Color: "#721412"},
	"ESXi":             {Name: "ESXi", Icon: iconLocal("esxi"), Color: "#607078"},
	"vSphere":          {Name: "vSphere", Icon: iconLocal("vsphere"), Color: "#4493f8"},
	"VMC":              {Name: "VMC", Icon: iconLocal("vmc"), Color: "#FF9900"},
	"Automation":       {Name: "Automation", Icon: iconDevicon("bash", "original"), Color: "#4EAA25"},
	"Multithreading":   {Name: "Multithreading", Icon: iconLocal("cplusplus"), Color: "#00599C"},
	"Storage":          {Name: "Storage", Icon: iconDevicon("mongodb", "original"), Color: "#47A248"},
	"Debugging":        {Name: "Debugging", Icon: iconDevicon("vscode", "original"), Color: "#007ACC"},
	"Scalability":      {Name: "Scalability", Icon: iconDevicon("apachekafka", "original"), Color: "#231F20"},
	"Ginkgo":           {Name: "Ginkgo", Icon: iconLocal("golang"), Color: "#00ADD8"},
	"Linux":            {Name: "Linux", Icon: iconLocal("linux"), Color: "#FCC624"},
	"Scripting":        {Name: "Scripting", Icon: iconDevicon("bash", "original"), Color: "#4EAA25"},
	"AWS Lambda":       {Name: "AWS Lambda", Icon: iconLocal("aws"), Color: "#FF9900"},
	"Cursor":           {Name: "Cursor", Icon: "", Color: "#A371F7"},
	"GitHub Copilot":   {Name: "GitHub Copilot", Icon: iconDevicon("github", "original"), Color: "#FFFFFF"},
	"Google Gemini":    {Name: "Google Gemini", Icon: iconDevicon("google", "original"), Color: "#4285F4"},
}

func iconLocal(name string) string {
	return "/static/icons/" + name + ".svg"
}

func iconDevicon(name, variant string) string {
	return "https://cdn.jsdelivr.net/gh/devicons/devicon@v2.16.0/icons/" + name + "/" + name + "-" + variant + ".svg"
}

func TechBadgeFor(name string) TechBadge {
	if badge, ok := techMeta[name]; ok {
		return badge
	}
	color := "#6366F1"
	return TechBadge{Name: name, Icon: "", Color: color}
}

func TechBadges(names []string) []TechBadge {
	badges := make([]TechBadge, len(names))
	for i, name := range names {
		badges[i] = TechBadgeFor(name)
	}
	return badges
}
