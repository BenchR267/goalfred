# Sometimes it's a README fix, or something like that - which isn't relevant for
# including in a project's CHANGELOG for example
declared_trivial = github.pr_title.include? '#trivial'

# Make it more obvious that a PR is a work in progress and shouldn't be merged yet
warn('PR is classed as Work in Progress') if github.pr_title.include? '[WIP]'

# Warn when there is a big PR
warn('Big PR (cc @kiliankoe @BenchR267)') if git.lines_of_code > 500

lint_output = `golint *.go`
if lint_output.length > 0
	fail("Please fix linting issues! \n\t#{lint_output}")
end

# Added changelog entry?
if !git.modified_files.include?('CHANGELOG.md') && !declared_trivial
  fail("Please include a CHANGELOG entry. \nYou can find it at [CHANGELOG.md](https://github.com/BenchR267/goalfred/blob/master/CHANGELOG.md).")
end
