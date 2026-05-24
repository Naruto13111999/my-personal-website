(function () {
  'use strict';

  const intro = document.getElementById('intro');
  const introEnter = document.getElementById('intro-enter');

  function enterSite() {
    if (!intro || intro.classList.contains('is-hidden')) return;
    intro.classList.add('is-hidden');
    document.body.classList.remove('intro-open');
    sessionStorage.setItem('gyanankur-intro-seen', '1');
    window.location.hash = 'home';
    document.getElementById('home')?.scrollIntoView({ behavior: 'smooth' });
  }

  if (intro && sessionStorage.getItem('gyanankur-intro-seen') === '1') {
    intro.classList.add('is-hidden');
    document.body.classList.remove('intro-open');
  } else if (introEnter) {
    introEnter.addEventListener('click', enterSite);
    document.addEventListener('keydown', (e) => {
      if (intro.classList.contains('is-hidden')) return;
      if (e.key === 'Enter' || e.key === ' ') {
        e.preventDefault();
        enterSite();
      }
    });
  }

  const nav = document.getElementById('nav');
  const navToggle = document.getElementById('nav-toggle');
  const navLinks = document.getElementById('nav-links');
  const navSections = ['about', 'experience', 'projects', 'skills', 'contact'];
  const navAnchors = navLinks?.querySelectorAll('a[href^="#"]') ?? [];

  function setActiveNav(sectionId) {
    navAnchors.forEach((link) => {
      const id = link.getAttribute('href')?.slice(1) ?? '';
      link.classList.toggle('nav-active', id === sectionId);
    });
  }

  function updateActiveNav() {
    const offset = nav.offsetHeight + 80;
    const scrollPos = window.scrollY + offset;
    let current = '';

    navSections.forEach((id) => {
      const section = document.getElementById(id);
      if (section && section.offsetTop <= scrollPos) {
        current = id;
      }
    });

    setActiveNav(current);
  }

  window.addEventListener('scroll', () => {
    nav.classList.toggle('scrolled', window.scrollY > 40);
    updateActiveNav();
  }, { passive: true });

  updateActiveNav();

  navToggle.addEventListener('click', () => {
    navLinks.classList.toggle('open');
  });

  navAnchors.forEach((link) => {
    link.addEventListener('click', () => {
      navLinks.classList.remove('open');
      const id = link.getAttribute('href')?.slice(1) ?? '';
      if (navSections.includes(id)) {
        setActiveNav(id);
      }
    });
  });

  const projectTabs = document.querySelectorAll('.projects-tab');
  const projectPanels = document.querySelectorAll('.company-projects[role="tabpanel"]');
  const projectsPrompt = document.getElementById('projects-prompt');

  function revealPanel(panel) {
    panel.classList.add('visible');
    panel.querySelectorAll('.project-row').forEach((row) => row.classList.add('visible'));
  }

  function showCompanyProjects(slug) {
    projectTabs.forEach((tab) => {
      const active = tab.dataset.company === slug;
      tab.classList.toggle('is-active', active);
      tab.setAttribute('aria-selected', active ? 'true' : 'false');
    });

    projectPanels.forEach((panel) => {
      const show = slug && panel.id === `projects-panel-${slug}`;
      panel.hidden = !show;
      if (show) {
        revealPanel(panel);
      }
    });

    projectsPrompt?.classList.toggle('is-hidden', Boolean(slug));
  }

  projectTabs.forEach((tab) => {
    tab.addEventListener('click', () => {
      showCompanyProjects(tab.dataset.company ?? '');
    });
  });

  showCompanyProjects('jumpcloud');

  const revealEls = document.querySelectorAll(
    '.section-title, .about-grid, .contrib-card, .timeline-item, .company-projects, .project-row, .tech-list-grid, .contact-card, .core-stack'
  );

  revealEls.forEach((el) => el.classList.add('reveal'));

  const observer = new IntersectionObserver(
    (entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          entry.target.classList.add('visible');
        }
      });
    },
    { threshold: 0.1, rootMargin: '0px 0px -40px 0px' }
  );

  revealEls.forEach((el) => observer.observe(el));
})();
