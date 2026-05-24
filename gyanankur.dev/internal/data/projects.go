package data

type ProjectGroup struct {
	Company  string    `json:"company"`
	Slug     string    `json:"slug"`
	Theme    string    `json:"theme"`
	Period   string    `json:"period"`
	Tagline  string    `json:"tagline"`
	AIUsage  string    `json:"ai_usage"`
	Projects []Project `json:"projects"`
}

func BuildProjectGroups(projects []Project) []ProjectGroup {
	meta := []struct {
		company, slug, theme, period, tagline, aiUsage string
	}{
		{
			company: "JumpCloud", slug: "jumpcloud", theme: "jc",
			period: "Dec 2024 – Present", tagline: "MDM platforms · SQS scalability · Go services",
			aiUsage: "Cursor, Gemini & Copilot for design, implementation, debugging, and reviews — accelerated localization delivery by 3 weeks.",
		},
		{
			company: "Symphony Fintech", slug: "symphony", theme: "sym",
			period: "Feb – Nov 2024", tagline: "Microsecond C++ · HFT market infrastructure",
			aiUsage: "AI-assisted Python tooling for instrument ID generation, log analysis, and faster iteration on market-feed optimizations.",
		},
		{
			company: "VMware", slug: "vmware", theme: "vm",
			period: "Aug 2022 – Jan 2024", tagline: "NFS core storage · ESXi · vSphere plugins",
			aiUsage: "AI-assisted Python scripts for nConnect testing, log triage, and faster P0/P1 NFS client debugging.",
		},
	}

	groups := make([]ProjectGroup, 0, len(meta))
	for _, m := range meta {
		var companyProjects []Project
		for _, p := range projects {
			if p.Company == m.company {
				companyProjects = append(companyProjects, p)
			}
		}
		if len(companyProjects) == 0 {
			continue
		}
		groups = append(groups, ProjectGroup{
			Company:  m.company,
			Slug:     m.slug,
			Theme:    m.theme,
			Period:   m.period,
			Tagline:  m.tagline,
			AIUsage:  m.aiUsage,
			Projects: companyProjects,
		})
	}
	return groups
}
