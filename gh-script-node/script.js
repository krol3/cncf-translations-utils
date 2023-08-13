// https://docs.github.com/en/rest/issues/issues?apiVersion=2022-11-28#create-an-issue

import { Octokit } from "@octokit/core";

const GITHUB_TOKEN = process.env.GITHUB_TOKEN

const octokit = new Octokit({
  auth: GITHUB_TOKEN
});

// const file_to_be_translated = "content/en/docs/concepts/architecture/cri.md"
const file_to_be_translated = "content/en/docs/" + process.argv[2]

const title = `[es] localize ${file_to_be_translated} to Spanish`

await octokit.request('POST /repos/kubernetes/website/issues', {
  owner: 'kubernetes',
  repo: 'website',
  title,
  body: `
**This is a Feature Request**

**What would you like to be added**
Localize \`${file_to_be_translated}\` to Spanish

**Why is this needed**
There is no Spanish localization for this file.

**Comments**

/triage accepted
/kind feature
/language es
/sig docs
  `,
  assignees: [],
  labels: [],
  headers: {
    'X-GitHub-Api-Version': '2022-11-28'
  }
})