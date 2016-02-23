package bintray

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/jfrogdev/jfrog-cli-go/bintray/commands"
	"github.com/jfrogdev/jfrog-cli-go/bintray/utils"
	"github.com/jfrogdev/jfrog-cli-go/cliutils"
	"os"
	"strconv"
	"strings"
)

func GetCommands() []cli.Command {
	return []cli.Command{
		{
			Name:    "config",
			Usage:   "Configure Bintray details",
			Aliases: []string{"c"},
			Flags:   getConfigFlags(),
			Action: func(c *cli.Context) {
				config(c)
			},
		},
		{
			Name:    "upload",
			Usage:   "Upload files",
			Aliases: []string{"u"},
			Flags:   getUploadFlags(),
			Action: func(c *cli.Context) {
				upload(c)
			},
		},
		{
			Name:    "download-file",
			Usage:   "Download file",
			Aliases: []string{"dlf"},
			Flags:   getDownloadFileFlags(),
			Action: func(c *cli.Context) {
				downloadFile(c)
			},
		},
		{
			Name:    "download-ver",
			Usage:   "Download Version files",
			Aliases: []string{"dlv"},
			Flags:   getDownloadVersionFlags(),
			Action: func(c *cli.Context) {
				downloadVersion(c)
			},
		},
		{
			Name:    "package-show",
			Usage:   "Show Package details",
			Aliases: []string{"ps"},
			Flags:   getFlags(),
			Action: func(c *cli.Context) {
				showPackage(c)
			},
		},
		{
			Name:    "package-create",
			Usage:   "Create Package",
			Aliases: []string{"pc"},
			Flags:   getCreateAndUpdatePackageFlags(),
			Action: func(c *cli.Context) {
				createPackage(c)
			},
		},
		{
			Name:    "package-update",
			Usage:   "Update Package",
			Aliases: []string{"pu"},
			Flags:   getCreateAndUpdatePackageFlags(),
			Action: func(c *cli.Context) {
				updatePackage(c)
			},
		},
		{
			Name:    "package-delete",
			Usage:   "Delete Package",
			Aliases: []string{"pd"},
			Flags:   getDeletePackageAndVersionFlags(),
			Action: func(c *cli.Context) {
				deletePackage(c)
			},
		},
		{
			Name:    "version-show",
			Usage:   "Show Version",
			Aliases: []string{"vs"},
			Flags:   getFlags(),
			Action: func(c *cli.Context) {
				showVersion(c)
			},
		},
		{
			Name:    "version-create",
			Usage:   "Create Version",
			Aliases: []string{"vc"},
			Flags:   getCreateAndUpdateVersionFlags(),
			Action: func(c *cli.Context) {
				createVersion(c)
			},
		},
		{
			Name:    "version-update",
			Usage:   "Update Version",
			Aliases: []string{"vu"},
			Flags:   getCreateAndUpdateVersionFlags(),
			Action: func(c *cli.Context) {
				updateVersion(c)
			},
		},
		{
			Name:    "version-delete",
			Usage:   "Delete Version",
			Aliases: []string{"vd"},
			Flags:   getDeletePackageAndVersionFlags(),
			Action: func(c *cli.Context) {
				deleteVersion(c)
			},
		},
		{
			Name:    "version-publish",
			Usage:   "Publish Version",
			Aliases: []string{"vp"},
			Flags:   getFlags(),
			Action: func(c *cli.Context) {
				publishVersion(c)
			},
		},
		{
			Name:    "entitlements",
			Usage:   "Manage Entitlements",
			Aliases: []string{"ent"},
			Flags:   getEntitlementsFlags(),
			Action: func(c *cli.Context) {
				entitlements(c)
			},
		},
		{
			Name:    "entitlement-keys",
			Usage:   "Manage Entitlement Keys",
			Aliases: []string{"ent-keys"},
			Flags:   getEntitlementKeysFlags(),
			Action: func(c *cli.Context) {
				entitlementKeys(c)
			},
		},
		{
			Name:    "url-sign",
			Usage:   "Create Signed Download URL",
			Aliases: []string{"us"},
			Flags:   getUrlSigningFlags(),
			Action: func(c *cli.Context) {
				signUrl(c)
			},
		},
		{
			Name:    "gpg-sign-file",
			Usage:   "GPF Sign file",
			Aliases: []string{"gsf"},
			Flags:   getGpgSigningFlags(),
			Action: func(c *cli.Context) {
				gpgSignFile(c)
			},
		},
		{
			Name:    "gpg-sign-ver",
			Usage:   "GPF Sign Version",
			Aliases: []string{"gsv"},
			Flags:   getGpgSigningFlags(),
			Action: func(c *cli.Context) {
				gpgSignVersion(c)
			},
		},
	}
}

func getFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "user",
			EnvVar: "BINTRAY_USER",
			Usage:  "[Optional] Bintray username. If not set, the subject sent as part of the command argument is used for authentication.",
		},
		cli.StringFlag{
			Name:   "key",
			EnvVar: "BINTRAY_KEY",
			Usage:  "[Mandatory] Bintray API key",
		},
	}
}

func getConfigFlags() []cli.Flag {
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "interactive",
			Usage: "[Default: true] Set to false if you do not want the config command to be interactive.",
		},
	}
	return append(flags, getFlags()...)
}

func getPackageFlags(prefix string) []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  prefix + "licenses",
			Value: "",
			Usage: "[Mandatory for OSS] Package licenses in the form of \"Apache-2.0\",\"GPL-3.0\"...",
		},
		cli.StringFlag{
			Name:  prefix + "vcs-url",
			Value: "",
			Usage: "[Mandatory for OSS] Package VCS URL.",
		},
		cli.StringFlag{
			Name:  prefix + "pub-dn",
			Value: "",
			Usage: "[Default: false] Public download numbers.",
		},
		cli.StringFlag{
			Name:  prefix + "pub-stats",
			Value: "",
			Usage: "[Default: true] Public statistics",
		},
		cli.StringFlag{
			Name:  prefix + "desc",
			Value: "",
			Usage: "[Optional] Package description.",
		},
		cli.StringFlag{
			Name:  prefix + "labels",
			Value: "",
			Usage: "[Optional] Package lables in the form of \"lable11\",\"lable2\"...",
		},
		cli.StringFlag{
			Name:  prefix + "cust-licenses",
			Value: "",
			Usage: "[Optional] Package custom licenses in the form of \"my-license-1\",\"my-license-2\"...",
		},
		cli.StringFlag{
			Name:  prefix + "website-url",
			Value: "",
			Usage: "[Optional] Package web site URL.",
		},
		cli.StringFlag{
			Name:  prefix + "issuetracker-url",
			Value: "",
			Usage: "[Optional] Package Issues Tracker URL.",
		},
		cli.StringFlag{
			Name:  prefix + "github-repo",
			Value: "",
			Usage: "[Optional] Package Github repository.",
		},
		cli.StringFlag{
			Name:  prefix + "github-rel-notes",
			Value: "",
			Usage: "[Optional] Github release notes file.",
		},
	}
}

func getVersionFlags(prefix string) []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  prefix + "github-tag-rel-notes",
			Value: "",
			Usage: "[Default: false] Set to true if you wish to use a Github tag release notes.",
		},
		cli.StringFlag{
			Name:  prefix + "desc",
			Value: "",
			Usage: "[Optional] Version description.",
		},
		cli.StringFlag{
			Name:  prefix + "released",
			Value: "",
			Usage: "[Optional] Release date in ISO8601 format (yyyy-MM-dd'T'HH:mm:ss.SSSZ)",
		},
		cli.StringFlag{
			Name:  prefix + "github-rel-notes",
			Value: "",
			Usage: "[Optional] Github release notes file.",
		},
		cli.StringFlag{
			Name:  prefix + "vcs-tag",
			Value: "",
			Usage: "[Optional] VCS tag.",
		},
	}
}

func getCreateAndUpdatePackageFlags() []cli.Flag {
	return append(getFlags(), getPackageFlags("")...)
}

func getCreateAndUpdateVersionFlags() []cli.Flag {
	return append(getFlags(), getVersionFlags("")...)
}

func getDeletePackageAndVersionFlags() []cli.Flag {
	return append(getFlags(), cli.StringFlag{
		Name:  "quiet",
		Value: "",
		Usage: "[Default: false] Set to true to skip the delete confirmation message.",
	})
}

func getDownloadFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "flat",
			Value: "",
			Usage: "[Default: false] Set to true if you do not wish to have the Bintray path structure created locally for your downloaded files.",
		},
		cli.StringFlag{
			Name:  "min-split",
			Value: "",
			Usage: "[Default: 5120] Minimum file size in KB to split into ranges when downloading. Set to -1 for no splits.",
		},
		cli.StringFlag{
			Name:  "split-count",
			Value: "",
			Usage: "[Default: 3] Number of parts to split a file when downloading. Set to 0 for no splits.",
		},
	}
}

func getDownloadFileFlags() []cli.Flag {
	return append(getFlags(), getDownloadFlags()...)
}

func getDownloadVersionFlags() []cli.Flag {
	flags := append(getFlags(), cli.StringFlag{
		Name:  "threads",
		Value: "",
		Usage: "[Default: 3] Number of artifacts to download in parallel.",
	})
	return append(flags, getDownloadFlags()...)
}

func getUploadFlags() []cli.Flag {
	flags := append(getFlags(), []cli.Flag{
		cli.StringFlag{
			Name:  "recursive",
			Value: "",
			Usage: "[Default: true] Set to false if you do not wish to collect files in sub-folders to be uploaded to Bintray.",
		},
		cli.StringFlag{
			Name:  "flat",
			Value: "",
			Usage: "[Default: true] If not set to true, and the upload path ends with a slash, files are uploaded according to their file system hierarchy.",
		},
		cli.BoolFlag{
			Name:  "regexp",
			Usage: "[Default: false] Set to true to use a regular expression instead of wildcards expression to collect files to upload.",
		},
		cli.StringFlag{
			Name:  "publish",
			Value: "",
			Usage: "[Default: false] Set to true to publish the uploaded files.",
		},
		cli.StringFlag{
			Name:  "override",
			Value: "",
			Usage: "[Default: false] Set to true to enable overriding existing published files.",
		},
		cli.StringFlag{
			Name:  "explode",
			Value: "",
			Usage: "[Default: false] Set to true to explode archived files after upload.",
		},
		cli.StringFlag{
			Name:  "threads",
			Value: "",
			Usage: "[Default: 3] Number of artifacts to upload in parallel.",
		},
		cli.BoolFlag{
			Name:  "dry-run",
			Usage: "[Default: false] Set to true to disable communication with Bintray.",
		},
	}...)
	flags = append(flags, getPackageFlags("pkg-")...)
	return append(flags, getVersionFlags("ver-")...)
}

func getEntitlementsFlags() []cli.Flag {
	return append(getFlags(), []cli.Flag{
		cli.StringFlag{
			Name:  "id",
			Usage: "[Optional] Entitlement ID. Used for Entitlements update.",
		},
		cli.StringFlag{
			Name:  "access",
			Usage: "[Optional] Entitlement access. Used for Entitlements creation and update.",
		},
		cli.StringFlag{
			Name:  "keys",
			Usage: "[Optional] Used for Entitlements creation and update. List of Download Keys in the form of \"key1\",\"key2\"...",
		},
		cli.StringFlag{
			Name:  "path",
			Usage: "[Optional] Entitlement path. Used for Entitlements creating and update.",
		},
	}...)
}

func getEntitlementKeysFlags() []cli.Flag {
	return append(getFlags(), []cli.Flag{
		cli.StringFlag{
			Name:  "org",
			Usage: "[Optional] Bintray organization",
		},
		cli.StringFlag{
			Name:  "id",
			Usage: "[Optional] Download Key ID (required for 'bt ent-keys show/create/update/delete'",
		},
		cli.StringFlag{
			Name:  "expiry",
			Usage: "[Optional] Download Key expiry (required for 'bt ent-keys show/create/update/delete'",
		},
		cli.StringFlag{
			Name:  "ex-check-url",
			Usage: "[Optional] Used for Download Key creation and update. You can optionally provide an existence check directive, in the form of a callback URL, to verify whether the source identity of the Download Key still exists.",
		},
		cli.StringFlag{
			Name:  "ex-check-cache",
			Usage: "[Optional] Used for Download Key creation and update. You can optionally provide the period in seconds for the callback URK results cache.",
		},
		cli.StringFlag{
			Name:  "white-cidrs",
			Usage: "[Optional] Used for Download Key creation and update. Specifying white CIDRs in the form of \"127.0.0.1/22\",\"193.5.0.1/92\" will allow access only for those IPs that exist in that address range.",
		},
		cli.StringFlag{
			Name:  "black-cidrs",
			Usage: "[Optional] Used for Download Key creation and update. Specifying black CIDRs in the foem of \"127.0.0.1/22\",\"193.5.0.1/92\" will block access for all IPs that exist in the specified range.",
		},
	}...)
}

func getUrlSigningFlags() []cli.Flag {
	return append(getFlags(), []cli.Flag{
		cli.StringFlag{
			Name:  "expiry",
			Usage: "[Optional] An expiry date for the URL, in Unix epoch time in milliseconds, after which the URL will be invalid. By default, expiry date will be 24 hours.",
		},
		cli.StringFlag{
			Name:  "valid-for",
			Usage: "[Optional] The number of seconds since generation before the URL expires. Mutually exclusive with the --expiry option.",
		},
		cli.StringFlag{
			Name:  "callback-id",
			Usage: "[Optional] An applicative identifier for the request. This identifier appears in download logs and is used in email and download webhook notifications.",
		},
		cli.StringFlag{
			Name:  "callback-email",
			Usage: "[Optional] An email address to send mail to when a user has used the download URL. This requiers a callback_id. The callback-id will be included in the mail message.",
		},
		cli.StringFlag{
			Name:  "callback-url",
			Usage: "[Optional] A webhook URL to call when a user has used the download URL.",
		},
		cli.StringFlag{
			Name:  "callback-method",
			Usage: "[Optional] HTTP method to use for making the callback. Will use POST by default. Supported methods are: GET, POST, PUT and HEAD.",
		},
	}...)
}

func getGpgSigningFlags() []cli.Flag {
	return append(getFlags(), cli.StringFlag{
		Name:  "passphrase",
		Usage: "[Optional] GPG passphrase.",
	})
}

func config(c *cli.Context) {
	if len(c.Args()) > 1 {
		cliutils.Exit(cliutils.ExitCodeError, "Wrong number of arguments. Try 'bt config --help'.")
	} else if len(c.Args()) == 1 {
		if c.Args()[0] == "show" {
			commands.ShowConfig()
		} else if c.Args()[0] == "clear" {
			commands.ClearConfig()
		} else {
			cliutils.Exit(cliutils.ExitCodeError, "Unknown argument '"+c.Args()[0]+"'. Available arguments are 'show' and 'clear'.")
		}
	} else {
		interactive := true
		if c.String("interactive") != "" {
			interactive = c.Bool("interactive")
		}
		if !interactive {
			if c.String("user") == "" || c.String("key") == "" {
				cliutils.Exit(cliutils.ExitCodeError, "The --url and --key options are mandatory when the --interactive option is set to false")
			}
		}
		bintrayDetails := createBintrayDetails(c, false)
		commands.Config(bintrayDetails, interactive)
	}
}

func showPackage(c *cli.Context) {
	if len(c.Args()) != 1 {
		cliutils.Exit(cliutils.ExitCodeError, "Wrong number of arguments. Try 'bt package-show --help'.")
	}
	packageDetails := utils.CreatePackageDetails(c.Args()[0])
	bintrayDetails := createBintrayDetails(c, true)
	commands.ShowPackage(packageDetails, bintrayDetails)
}

func showVersion(c *cli.Context) {
	if len(c.Args()) != 1 {
		cliutils.Exit(cliutils.ExitCodeError, "Wrong number of arguments. Try 'bt version-show --help'.")
	}
	versionDetails := utils.CreateVersionDetails(c.Args()[0])
	bintrayDetails := createBintrayDetails(c, true)
	commands.ShowVersion(versionDetails, bintrayDetails)
}

func createPackage(c *cli.Context) {
	if len(c.Args()) != 1 {
		cliutils.Exit(cliutils.ExitCodeError, "Wrong number of arguments. Try 'bt package-create --help'.")
	}
	packageDetails := utils.CreatePackageDetails(c.Args()[0])
	packageFlags := createPackageFlags(c, "")
	commands.CreatePackage(packageDetails, packageFlags)
}

func createVersion(c *cli.Context) {
	if len(c.Args()) != 1 {
		cliutils.Exit(cliutils.ExitCodeError, "Wrong number of arguments. Try 'bt version-create --help'.")
	}
	versionDetails := utils.CreateVersionDetails(c.Args()[0])
	versionFlags := createVersionFlags(c, "")
	commands.CreateVersion(versionDetails, versionFlags)
}

func updateVersion(c *cli.Context) {
	if len(c.Args()) != 1 {
		cliutils.Exit(cliutils.ExitCodeError, "Wrong number of arguments. Try 'bt version-update --help'.")
	}
	versionDetails := utils.CreateVersionDetails(c.Args()[0])
	versionFlags := createVersionFlags(c, "")
	commands.UpdateVersion(versionDetails, versionFlags)
}

func updatePackage(c *cli.Context) {
	if len(c.Args()) != 1 {
		cliutils.Exit(cliutils.ExitCodeError, "Wrong number of arguments. Try 'bt package-update --help'.")
	}
	packageDetails := utils.CreatePackageDetails(c.Args()[0])
	packageFlags := createPackageFlags(c, "")
	commands.UpdatePackage(packageDetails, packageFlags)
}

func deletePackage(c *cli.Context) {
	if len(c.Args()) != 1 {
		cliutils.Exit(cliutils.ExitCodeError, "Wrong number of arguments. Try 'bt package-delete --help'.")
	}
	packageDetails := utils.CreatePackageDetails(c.Args()[0])
	bintrayDetails := createBintrayDetails(c, true)

	if !c.Bool("quiet") {
		var confirm string
		fmt.Print("Delete package " + packageDetails.Package + "? (y/n): ")
		fmt.Scanln(&confirm)
		if !cliutils.ConfirmAnswer(confirm) {
			return
		}
	}
	commands.DeletePackage(packageDetails, bintrayDetails)
}

func deleteVersion(c *cli.Context) {
	if len(c.Args()) != 1 {
		cliutils.Exit(cliutils.ExitCodeError, "Wrong number of arguments. Try 'bt version-delete --help'.")
	}
	versionDetails := utils.CreateVersionDetails(c.Args()[0])
	bintrayDetails := createBintrayDetails(c, true)

	if !c.Bool("quiet") {
		var confirm string
		fmt.Print("Delete version " + versionDetails.Version +
			" of package " + versionDetails.Package + "? (y/n): ")
		fmt.Scanln(&confirm)
		if !cliutils.ConfirmAnswer(confirm) {
			return
		}
	}
	commands.DeleteVersion(versionDetails, bintrayDetails)
}

func publishVersion(c *cli.Context) {
	if len(c.Args()) != 1 {
		cliutils.Exit(cliutils.ExitCodeError, "Wrong number of arguments. Try 'bt version-publish --help'.")
	}
	versionDetails := utils.CreateVersionDetails(c.Args()[0])
	bintrayDetails := createBintrayDetails(c, true)
	commands.PublishVersion(versionDetails, bintrayDetails)
}

func downloadVersion(c *cli.Context) {
	if len(c.Args()) != 1 {
		cliutils.Exit(cliutils.ExitCodeError, "Wrong number of arguments. Try 'bt download-ver --help'.")
	}
	versionDetails := commands.CreateVersionDetailsForDownloadVersion(c.Args()[0])
	flags := createDownloadFlags(c)
	commands.DownloadVersion(versionDetails, flags)
}

func upload(c *cli.Context) {
	if len(c.Args()) != 2 {
		cliutils.Exit(cliutils.ExitCodeError, "Wrong number of arguments. Try 'bt upload --help'.")
	}
	localPath := c.Args()[0]
	versionDetails, uploadPath := utils.CreateVersionDetailsAndPath(c.Args()[1])
	uploadFlags := createUploadFlags(c)
	packageFlags := createPackageFlags(c, "pkg-")
	versionFlags := createVersionFlags(c, "ver-")
	commands.Upload(versionDetails, localPath, uploadPath, uploadFlags, packageFlags, versionFlags)
}

func downloadFile(c *cli.Context) {
	if len(c.Args()) != 1 {
		cliutils.Exit(cliutils.ExitCodeError, "Wrong number of arguments. Try 'bt download-file --help'.")
	}
	versionDetails, path := utils.CreatePackageDetailsAndPath(c.Args()[0])
	flags := createDownloadFlags(c)
	commands.DownloadFile(versionDetails, path, flags)
}

func signUrl(c *cli.Context) {
	if len(c.Args()) != 1 {
		cliutils.Exit(cliutils.ExitCodeError, "Wrong number of arguments. Try 'bt url-sign --help'.")
	}
	urlSigningDetails := utils.CreatePathDetails(c.Args()[0])
	urlSigningFlags := createUrlSigningFlags(c)
	commands.SignVersion(urlSigningDetails, urlSigningFlags)
}

func gpgSignFile(c *cli.Context) {
	if len(c.Args()) != 1 {
		cliutils.Exit(cliutils.ExitCodeError, "Wrong number of arguments. Try 'bt url-sign --help'.")
	}
	pathDetails := utils.CreatePathDetails(c.Args()[0])
	commands.GpgSignFile(pathDetails, c.String("passphrase"), createBintrayDetails(c, true))
}

func gpgSignVersion(c *cli.Context) {
	if len(c.Args()) != 1 {
		cliutils.Exit(cliutils.ExitCodeError, "Wrong number of arguments. Try 'bt url-sign --help'.")
	}
	versionDetails := utils.CreateVersionDetails(c.Args()[0])
	commands.GpgSignVersion(versionDetails, c.String("passphrase"), createBintrayDetails(c, true))
}

func entitlementKeys(c *cli.Context) {
	org := c.String("org")
	argsSize := len(c.Args())
	if argsSize == 0 {
		bintrayDetails := createBintrayDetails(c, true)
		commands.ShowDownloadKeys(bintrayDetails, org)
		return
	}
	if argsSize != 2 {
		cliutils.Exit(cliutils.ExitCodeError, "Wrong number of arguments. Try 'bt ent-keys --help'.")
	}
	keyId := c.Args()[1]
	if c.Args()[0] == "show" {
		commands.ShowDownloadKey(createDownloadKeyFlagsForShowAndDelete(keyId, c), org)
	} else if c.Args()[0] == "create" {
		commands.CreateDownloadKey(createDownloadKeyFlagsForCreateAndUpdate(keyId, c), org)
	} else if c.Args()[0] == "update" {
		commands.UpdateDownloadKey(createDownloadKeyFlagsForCreateAndUpdate(keyId, c), org)
	} else if c.Args()[0] == "delete" {
		commands.DeleteDownloadKey(createDownloadKeyFlagsForShowAndDelete(keyId, c), org)
	} else {
		cliutils.Exit(cliutils.ExitCodeError, "Expecting show, create, update or delete after the key argument. Got "+c.Args()[0])
	}
}

func entitlements(c *cli.Context) {
	argsSize := len(c.Args())
	if argsSize == 0 {
		cliutils.Exit(cliutils.ExitCodeError, "Wrong number of arguments. Try 'bt ent --help'.")
	}
	if argsSize == 1 {
		bintrayDetails := createBintrayDetails(c, true)
		details := commands.CreateVersionDetailsForEntitlements(c.Args()[0])
		commands.ShowEntitlements(bintrayDetails, details)
		return
	}
	if argsSize != 2 {
		cliutils.Exit(cliutils.ExitCodeError, "Wrong number of arguments. Try 'bt ent --help'.")
	}
	details := commands.CreateVersionDetailsForEntitlements(c.Args()[1])
	if c.Args()[0] == "show" {
		commands.ShowEntitlement(createEntitlementFlagsForShowAndDelete(c), details)
	} else if c.Args()[0] == "create" {
		commands.CreateEntitlement(createEntitlementFlagsForCreate(c), details)
	} else if c.Args()[0] == "update" {
		commands.UpdateEntitlement(createEntitlementFlagsForUpdate(c), details)
	} else if c.Args()[0] == "delete" {
		commands.DeleteEntitlement(createEntitlementFlagsForShowAndDelete(c), details)
	} else {
		cliutils.Exit(cliutils.ExitCodeError, "Expecting show, create, update or delete before "+c.Args()[1]+". Got "+c.Args()[0])
	}
}

func createPackageFlags(c *cli.Context, prefix string) *utils.PackageFlags {
	var publicDownloadNumbers string
	var publicStats string
	if c.String(prefix+"pub-dn") != "" {
		publicDownloadNumbers = c.String(prefix + "pub-dn")
		publicDownloadNumbers = strings.ToLower(publicDownloadNumbers)
		if publicDownloadNumbers != "true" && publicDownloadNumbers != "false" {
			cliutils.Exit(cliutils.ExitCodeError, "The --"+prefix+"pub-dn option should have a boolean value.")
		}
	}
	if c.String(prefix+"pub-stats") != "" {
		publicStats = c.String(prefix + "pub-stats")
		publicStats = strings.ToLower(publicStats)
		if publicStats != "true" && publicStats != "false" {
			cliutils.Exit(cliutils.ExitCodeError, "The --"+prefix+"pub-stats option should have a boolean value.")
		}
	}

	return &utils.PackageFlags{
		BintrayDetails:         createBintrayDetails(c, true),
		Desc:                   c.String(prefix + "desc"),
		Labels:                 c.String(prefix + "labels"),
		Licenses:               c.String(prefix + "licenses"),
		CustomLicenses:         c.String(prefix + "cust-licenses"),
		VcsUrl:                 c.String(prefix + "vcs-url"),
		WebsiteUrl:             c.String(prefix + "website-url"),
		IssueTrackerUrl:        c.String(prefix + "issuetracker-url"),
		GithubRepo:             c.String(prefix + "github-repo"),
		GithubReleaseNotesFile: c.String(prefix + "github-rel-notes"),
		PublicDownloadNumbers:  publicDownloadNumbers,
		PublicStats:            publicStats}
}

func createVersionFlags(c *cli.Context, prefix string) *utils.VersionFlags {
	var githubTagReleaseNotes string
	if c.String(prefix+"github-tag-rel-notes") != "" {
		githubTagReleaseNotes = c.String(prefix + "github-tag-rel-notes")
		githubTagReleaseNotes = strings.ToLower(githubTagReleaseNotes)
		if githubTagReleaseNotes != "true" && githubTagReleaseNotes != "false" {
			cliutils.Exit(cliutils.ExitCodeError, "The --github-tag-rel-notes option should have a boolean value.")
		}
	}
	return &utils.VersionFlags{
		BintrayDetails:           createBintrayDetails(c, true),
		Desc:                     c.String(prefix + "desc"),
		VcsTag:                   c.String(prefix + "vcs-tag"),
		Released:                 c.String(prefix + "released"),
		GithubReleaseNotesFile:   c.String(prefix + "github-rel-notes"),
		GithubUseTagReleaseNotes: githubTagReleaseNotes}
}

func createUrlSigningFlags(c *cli.Context) *commands.UrlSigningFlags {
	if c.String("valid-for") != "" {
		_, err := strconv.ParseInt(c.String("valid-for"), 10, 64)
		if err != nil {
			cliutils.Exit(cliutils.ExitCodeError, "The '--valid-for' option should have a numeric value.")
		}
	}

	return &commands.UrlSigningFlags{
		BintrayDetails: createBintrayDetails(c, true),
		Expiry:         c.String("expiry"),
		ValidFor:       c.String("valid-for"),
		CallbackId:     c.String("callback-id"),
		CallbackEmail:  c.String("callback-email"),
		CallbackUrl:    c.String("callback-url"),
		CallbackMethod: c.String("callback-method")}
}

func createUploadFlags(c *cli.Context) *commands.UploadFlags {
	return &commands.UploadFlags{
		BintrayDetails: createBintrayDetails(c, true),
		Recursive:      getBoolFlagValue(c, "recursive", true),
		Flat:           getBoolFlagValue(c, "flat", true),
        Publish:        getBoolFlagValue(c, "publish", false),
        Override:       getBoolFlagValue(c, "override", false),
        Explode:        getBoolFlagValue(c, "explode", false),
		UseRegExp:      c.Bool("regexp"),
		Threads:        getThreadsOptionValue(c),
		DryRun:         c.Bool("dry-run")}
}

func getThreadsOptionValue(c *cli.Context) (threads int) {
	if c.String("threads") == "" {
		threads = 3
	} else {
		var err error
		threads, err = strconv.Atoi(c.String("threads"))
		if err != nil || threads < 1 {
			cliutils.Exit(cliutils.ExitCodeError, "The '--threads' option should have a numeric positive value.")
		}
	}
	return
}

func createDownloadFlags(c *cli.Context) *utils.DownloadFlags {
	flat := false
	if c.String("flat") != "" {
		flat = c.Bool("flat")
	}
	return &utils.DownloadFlags{
		BintrayDetails: createBintrayDetails(c, true),
		Threads:        getThreadsOptionValue(c),
		MinSplitSize:   getMinSplitFlag(c),
		SplitCount:     getSplitCountFlag(c),
		Flat:           flat}
}

func createEntitlementFlagsForShowAndDelete(c *cli.Context) *commands.EntitlementFlags {
	if c.String("id") == "" {
		cliutils.Exit(cliutils.ExitCodeError, "Please add the --id option")
	}
	return &commands.EntitlementFlags{
		BintrayDetails: createBintrayDetails(c, true),
		Id:             c.String("id")}
}

func createEntitlementFlagsForCreate(c *cli.Context) *commands.EntitlementFlags {
	if c.String("access") == "" {
		cliutils.Exit(cliutils.ExitCodeError, "Please add the --access option")
	}
	return &commands.EntitlementFlags{
		BintrayDetails: createBintrayDetails(c, true),
		Path:           c.String("path"),
		Access:         c.String("access"),
		Keys:           c.String("keys")}
}

func createEntitlementFlagsForUpdate(c *cli.Context) *commands.EntitlementFlags {
	if c.String("id") == "" {
		cliutils.Exit(cliutils.ExitCodeError, "Please add the --id option")
	}
	if c.String("access") == "" {
		cliutils.Exit(cliutils.ExitCodeError, "Please add the --access option")
	}
	return &commands.EntitlementFlags{
		BintrayDetails: createBintrayDetails(c, true),
		Id:             c.String("id"),
		Path:           c.String("path"),
		Access:         c.String("access"),
		Keys:           c.String("keys")}
}

func createDownloadKeyFlagsForShowAndDelete(keyId string, c *cli.Context) *commands.DownloadKeyFlags {
	return &commands.DownloadKeyFlags{
		BintrayDetails: createBintrayDetails(c, true),
		Id:             keyId}
}

func createDownloadKeyFlagsForCreateAndUpdate(keyId string, c *cli.Context) *commands.DownloadKeyFlags {
	var cachePeriod int
	if c.String("ex-check-cache") != "" {
		var err error
		cachePeriod, err = strconv.Atoi(c.String("ex-check-cache"))
		if err != nil {
			cliutils.Exit(cliutils.ExitCodeError, "The --ex-check-cache option should have a numeric value.")
		}
	}
	return &commands.DownloadKeyFlags{
		BintrayDetails:      createBintrayDetails(c, true),
		Id:                  keyId,
		Expiry:              c.String("expiry"),
		ExistenceCheckUrl:   c.String("ex-check-url"),
		ExistenceCheckCache: cachePeriod,
		WhiteCidrs:          c.String("white-cidrs"),
		BlackCidrs:          c.String("black-cidrs")}
}

func createBintrayDetails(c *cli.Context, includeConfig bool) *cliutils.BintrayDetails {
	user := c.String("user")
	key := c.String("key")
	if includeConfig && (user == "" || key == "") {
		confDetails := commands.GetConfig()
		if user == "" {
			user = confDetails.User
		}
		if key == "" {
			key = confDetails.Key
		}
		if key == "" {
			cliutils.Exit(cliutils.ExitCodeError, "Please set your Bintray API key using config command, --key option or the BINTRAY_KEY envrionemt variable")
		}
	}
	apiUrl := os.Getenv("BINTRAY_API_URL")
	if apiUrl == "" {
		apiUrl = "https://api.bintray.com/"
	}
	downloadServerUrl := os.Getenv("BINTRAY_DOWNLOAD_URL")
	if downloadServerUrl == "" {
		downloadServerUrl = "https://dl.bintray.com/"
	}
	apiUrl = cliutils.AddTrailingSlashIfNeeded(apiUrl)
	downloadServerUrl = cliutils.AddTrailingSlashIfNeeded(downloadServerUrl)
	return &cliutils.BintrayDetails{
		ApiUrl:            apiUrl,
		DownloadServerUrl: downloadServerUrl,
		User:              user,
		Key:               key}
}

func getMinSplitFlag(c *cli.Context) int64 {
	if c.String("min-split") == "" {
		return 5120
	}
	minSplit, err := strconv.ParseInt(c.String("min-split"), 10, 64)
	if err != nil {
		cliutils.Exit(cliutils.ExitCodeError, "The '--min-split' option should have a numeric value. Try 'art download --help'.")
	}
	return minSplit
}

func getSplitCountFlag(c *cli.Context) int {
	if c.String("split-count") == "" {
		return 3
	}
	splitCount, err := strconv.Atoi(c.String("split-count"))
	if err != nil {
		cliutils.Exit(cliutils.ExitCodeError, "The '--split-count' option should have a numeric value. Try 'art download --help'.")
	}
	if splitCount > 15 {
		cliutils.Exit(cliutils.ExitCodeError, "The '--split-count' option value is limitted to a maximum of 15.")
	}
	if splitCount < 0 {
		cliutils.Exit(cliutils.ExitCodeError, "The '--split-count' option cannot have a negative value.")
	}
	return splitCount
}

func getBoolFlagValue(c *cli.Context, flagName string, defValue bool) bool {
	if c.String(flagName) == "" {
		return defValue
	}
    return c.Bool(flagName)
}