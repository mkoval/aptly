package cmd

import (
	"fmt"
	"github.com/smira/aptly/deb"
	"github.com/smira/aptly/query"
	"github.com/smira/commander"
	"github.com/smira/flag"
)

func aptlyPackageSearch(cmd *commander.Command, args []string) error {
	var err error
	if len(args) != 1 {
		cmd.Usage()
		return commander.ErrCommandError
	}

	q, err := query.Parse(args[0])
	if err != nil {
		return fmt.Errorf("unable to search: %s", err)
	}

	result := q.Query(context.CollectionFactory().PackageCollection())
	result.ForEach(func(p *deb.Package) error {
		context.Progress().Printf("%s\n", p)
		return nil
	})

	return err
}

func makeCmdPackageSearch() *commander.Command {
	cmd := &commander.Command{
		Run:       aptlyPackageSearch,
		UsageLine: "search <package-query>",
		Short:     "search for packages matching query",
		Long: `
Command search displays list of packages in whole DB that match package query

Example:

    $ aptly package search '$Architecture (i386), Name (% *-dev)'
`,
		Flag: *flag.NewFlagSet("aptly-package-search", flag.ExitOnError),
	}

	return cmd
}
