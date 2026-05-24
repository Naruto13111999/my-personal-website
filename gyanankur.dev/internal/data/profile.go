package data

type Profile struct {
	Name             string         `json:"name"`
	Title            string         `json:"title"`
	Tagline          string         `json:"tagline"`
	Email            string         `json:"email"`
	Phone            string         `json:"phone"`
	Location         string         `json:"location"`
	CoreStack        []string       `json:"core_stack"`
	Skills           []string       `json:"skills"`
	Education        Education      `json:"education"`
	Experience       []Experience   `json:"experience"`
	Projects         []Project      `json:"projects"`
	ProjectGroups    []ProjectGroup `json:"project_groups"`
	Activity         ActivityGraph  `json:"activity"`
}

type Education struct {
	Degree string `json:"degree"`
	School string `json:"school"`
	Period string `json:"period"`
}

type Experience struct {
	Role       string   `json:"role"`
	Company    string   `json:"company"`
	Period     string   `json:"period"`
	Highlights []string `json:"highlights"`
}

type Project struct {
	Title       string   `json:"title"`
	Company     string   `json:"company"`
	Type        string   `json:"type"`
	Highlight   string   `json:"highlight"`
	Description string   `json:"description"`
	AIUsage     string   `json:"ai_usage"`
	Tech        []string `json:"tech"`
	Impact      string   `json:"impact"`
	Featured    bool     `json:"featured"`
}

func GetProfile() Profile {
	projects := []Project{
		{
			Title: "Dynamic Queue Partitioning", Company: "JumpCloud", Type: "New Feature", Featured: true,
			Highlight:   "One org's 120K SQS backlog was freezing every other partition",
			Description: "Noisy-neighbor scalability for Amazon SQS — when one organization's partition backlog exceeded 120K in-flight messages, it blocked processing for all other partitions. Built dynamic queue partitioning with automatic scaling and intelligent load distribution.",
			Tech:        []string{"Golang", "Amazon SQS", "gRPC", "PostgreSQL"},
			Impact:      "Throughput, fault isolation, and resilience under enterprise load.",
		},
		{
			Title: "Diversion Queue", Company: "JumpCloud", Type: "New Feature",
			Highlight:   "Isolate noisy tenants before they jam the whole queue",
			Description: "End-to-end diversion pipeline that routes high-volume organizations away from shared partitions, keeping critical message paths healthy during traffic spikes.",
			Tech:        []string{"Golang", "gRPC", "PostgreSQL", "Amazon SQS"},
			Impact:      "Reliable throughput when individual orgs spike traffic.",
		},
		{
			Title: "Unified Localization Engine", Company: "JumpCloud", Type: "New Feature",
			Highlight:   "Cyrillic Windows fleets — shipped 3 weeks ahead of schedule",
			Description: "Backend services for localization and language handling on Cyrillic-based Windows devices. Owned design through deployment with AI-assisted iteration.",
			AIUsage:     "Cursor & Gemini for API design, boilerplate, and test scaffolding — cut delivery time by 3 weeks.",
			Tech:        []string{"Golang", "gRPC", "Windows MDM"},
			Impact:      "Accelerated delivery with agentic AI tooling.",
		},
		{
			Title: "iOS & iPad Wallpaper Policy", Company: "JumpCloud", Type: "New Feature",
			Highlight:   "Fleet-wide wallpaper control for Apple devices",
			Description: "gRPC endpoints, PostgreSQL persistence, and object storage for managed wallpaper policies on iOS and iPadOS — full lifecycle from design to production.",
			Tech:        []string{"Golang", "gRPC", "PostgreSQL", "Object Storage"},
			Impact:      "Enterprise wallpaper management at Apple fleet scale.",
		},
		{
			Title: "Delivery Optimization Policy", Company: "JumpCloud", Type: "New Feature",
			Highlight:   "Bandwidth wins for our top 3 enterprise customers",
			Description: "Windows policy backend that reduces content delivery latency and bandwidth consumption across managed endpoints.",
			Tech:        []string{"Golang", "Windows Policy", "gRPC"},
			Impact:      "Directly addressed top-customer performance concerns.",
		},
		{
			Title: "Device Policy Backend (SCEP & Taskbar)", Company: "JumpCloud", Type: "New Feature",
			Highlight:   "15% device adoption on the first production release",
			Description: "SCEP certificate enrollment and Windows Taskbar policy backends, including SCEP Extended Key Usage support for enterprise security.",
			Tech:        []string{"Golang", "SCEP", "gRPC", "PostgreSQL"},
			Impact:      "15% device adoption in first release.",
		},
		{
			Title: "Android App Version Tracking", Company: "JumpCloud", Type: "Enhancement",
			Highlight:   "Compliance visibility for managed Android fleets",
			Description: "Real-time version monitoring, compliance auditing, and exportable insights for enterprise Android app deployments.",
			Tech:        []string{"Golang", "PostgreSQL", "Android EMM"},
			Impact:      "Audit-ready compliance across Android estates.",
		},
		{
			Title: "Lost Mode (iOS & Android)", Company: "JumpCloud", Type: "New Feature", Featured: true,
			Highlight:   "Lock, message & locate lost fleet devices in real time",
			Description: "Remote lock, custom on-screen messaging, and location tracking for iOS (MDM) and Android (EMM). Full backend with Ginkgo unit tests and Python E2E coverage.",
			Tech:        []string{"Golang", "gRPC", "PostgreSQL", "Ginkgo", "Python"},
			Impact:      "Critical security capability for device loss scenarios.",
		},
		{
			Title: "Orderbook Snapshot Recovery", Company: "Symphony Fintech", Type: "New Feature",
			Highlight:   "Recover orderbook state after market disruptions",
			Description: "C++ TBT (Tick-by-Tick) snapshot recovery using Solace messaging — restores trading orderbooks after feed interruptions.",
			Tech:        []string{"C++", "Solace", "HFT", "Low Latency"},
			Impact:      "Core reliability for HFT infrastructure.",
		},
		{
			Title: "Market Data Compression Engine", Company: "Symphony Fintech", Type: "Enhancement", Featured: true,
			Highlight:   "Eliminated 15% CPU spent on market-feed compression",
			Description: "Moved per-session compression into C++ at the subscription layer, removing a bottleneck that lived in upper market-feed tiers.",
			Tech:        []string{"C++", "Multithreading", "Market Feed"},
			Impact:      "15% CPU reclaimed on hot paths.",
		},
		{
			Title: "Touchline Packet Generator", Company: "Symphony Fintech", Type: "Enhancement",
			Highlight:   "Faster packets by moving work closer to the wire",
			Description: "Generated Touchline packets directly from market feed data in C++, offloading the Market Feed Manager upper layer.",
			Tech:        []string{"C++", "Market Feed", "Low Latency"},
			Impact:      "Lower latency on subscription-based transmission.",
		},
		{
			Title: "Exchange Instrument ID Generator", Company: "Symphony Fintech", Type: "New Feature",
			Highlight:   "10× faster instrument subscriptions",
			Description: "Python tooling to derive Exchange Instrument IDs from master files, streamlining HFT subscription workflows.",
			AIUsage:     "Copilot-assisted Python for parsing master files and subscription automation.",
			Tech:        []string{"Python", "HFT", "Automation"},
			Impact:      "10× subscription performance improvement.",
		},
		{
			Title: "NFS vSphere Configuration Profiles", Company: "VMware", Type: "New Feature",
			Highlight:   "Declarative NFS config for vSphere at scale",
			Description: "Cross-BU NFS plugins for vSphere Configuration Profiles — declarative storage configuration for enterprise ESXi fleets.",
			Tech:        []string{"C", "C++", "NFS", "vSphere", "ESXi"},
			Impact:      "NFS workflows integrated into vSphere config management.",
		},
		{
			Title: "NFSv4.1 Client Control Path Optimization", Company: "VMware", Type: "Enhancement",
			Highlight:   "Faster NFS control path & shorter P0/P1 debug cycles",
			Description: "ESXi NFSv4.1 client control-path performance work plus Python tooling for nConnect bandwidth tests and log triage.",
			AIUsage:     "Gemini & Copilot for log triage scripts and nConnect test automation.",
			Tech:        []string{"C", "C++", "Python", "NFS", "TCP/IP"},
			Impact:      "Better bandwidth utilization and faster incident response.",
		},
		{
			Title: "VMC Event Automation", Company: "VMware", Type: "Enhancement",
			Highlight:   "Event-driven ops automation for VMware Cloud",
			Description: "AWS Lambda handlers for operational tasks triggered by events inside VMware Cloud (VMC) environments.",
			Tech:        []string{"AWS Lambda", "Python", "VMC"},
			Impact:      "Less manual toil for cloud infrastructure ops.",
		},
	}

	return Profile{
		Name:     "Gyanankur Dey",
		Title:    "Software Development Engineer",
		Tagline:  "Building scalable backends with AI-augmented engineering — low-latency systems and enterprise platforms.",
		Email:    "ankurdey429@gmail.com",
		Phone:    "+91-9330391621",
		Location: "India",
		CoreStack: []string{
			"Golang", "gRPC", "PostgreSQL", "AWS", "Docker", "C++",
		},
		Skills: []string{
			"Golang", "gRPC", "PostgreSQL", "Redis", "Amazon SQS",
			"C/C++", "Python", "Protocol Buffers", "Docker", "Kubernetes",
			"Linux", "Scripting", "AWS", "CI/CD", "REST API",
			"Cursor", "GitHub Copilot", "Google Gemini",
			"Multithreading", "TCP/IP", "NFS", "Storage",
			"Data Structures & Algorithms", "Debugging", "Scalability",
		},
		Education: Education{
			Degree: "B.Tech in Computer Science & Engineering",
			School: "Institute of Engineering & Management (IEM)",
			Period: "Aug 2018 – Aug 2022",
		},
		Experience: []Experience{
			{
				Role:    "Software Development Engineer 2",
				Company: "JumpCloud",
				Period:  "Dec 2024 – Present",
				Highlights: []string{
					"Designing scalable device policy backends in Go for enterprise MDM/EMM.",
					"Leading production support for P0–P2 incidents and system reliability.",
					"Leveraging AI-assisted engineering to accelerate delivery and code quality.",
				},
			},
			{
				Role:    "Software Development Engineer",
				Company: "Symphony Fintech",
				Period:  "Feb 2024 – Nov 2024",
				Highlights: []string{
					"Low-latency C++ development for high-frequency trading systems.",
					"Built orderbook recovery, compression, and market feed optimizations.",
					"Improved subscription performance 10× with AI-assisted Python automation.",
				},
			},
			{
				Role:    "Software Development Engineer",
				Company: "VMware",
				Period:  "Aug 2022 – Jan 2024",
				Highlights: []string{
					"Core Storage (NFS) development on ESXi in C/C++.",
					"Implemented NFS plugins for vSphere Configuration Profiles.",
					"AI-assisted Python scripts for nConnect testing and faster P0/P1 debug cycles.",
				},
			},
		},
		Projects:      projects,
		ProjectGroups: BuildProjectGroups(projects),
		Activity:         GetActivityGraph(),
	}
}
