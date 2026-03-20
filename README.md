# Resume
My resume, rendered using html, css and golang

## How to use

1. Create a file `data.yaml` and fill your data. Use `data.sample.yaml` for reference

2. Run `make`

3. It will generate html files based on the templates. Currently there are three templates, `euro`, `general`, and `portfolio`. You can edit them in `templates` folder or add your own.

4. Open the generated html file in your browser. Print the page as pdf to get pdf version of the resume.

## GitHub Pages

This repo includes a workflow at `.github/workflows/deploy-pages.yml` that publishes the generated portfolio to GitHub Pages.

1. In GitHub repo settings, set Pages source to `GitHub Actions`
2. Push to `main` or run the workflow manually
3. The workflow publishes `portfolio.html` as the site `index.html`
