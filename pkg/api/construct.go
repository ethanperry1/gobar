package api

type commandArguments struct {
	argType ArgType
	value   string
}

func (arg *commandArguments) Type() ArgType {
	return arg.argType
}

func (arg *commandArguments) Value() string {
	return arg.value
}

type commandSurface struct {
	parent CommandSurface
	command CommandArgument
}

func (surface *commandSurface) Parent() (CommandSurface, bool) {
	return surface.parent, surface.parent != nil
}

func (surface *commandSurface) Command() CommandArgument {
	return surface.command
}

type packagesRuleBuilder struct{
	CommandSurface
}

func (builder *packagesRuleBuilder) Filter(s string) PackageContext {
	return &packageContextRuleBuilder{
		CommandSurface: &commandSurface{
			parent:  builder.CommandSurface,
			command: &commandArguments{
				argType: Fltr,
				value: s,
			},
		},
	}
}

type packageContextRuleBuilder struct {
	CommandSurface
}

func (builder *packageContextRuleBuilder) Files() Files {
	return &filesRuleBuilder{
		CommandSurface: &commandSurface{
			
		}
	}
}

type filesRuleBuilder struct{
	CommandSurface
}

func (builder *filesRuleBuilder) Filter(s string) FileContext {
	return &fileContextRuleBuilder{
		filter: s,
	}
}

type fileContextRuleBuilder struct {
	filter string
}

func (builder *fileContextRuleBuilder) Functions() Functions {
	return &functionsRuleBuilder{}
}

type functionsRuleBuilder struct {
}

func (builder *functionsRuleBuilder) Filter(s string) CommandSurface {
	return &functionCommandSurface{}
}

type functionCommandSurface struct {
}

func AllPackages() Packages {
	return &packagesRuleBuilder{
		CommandSurface: &commandSurface{
			parent: nil,
			command: &commandArguments{
				argType: Any,
			},
		},
	}
}

func Min(float64, ...CommandSurface) Option {

	return func() {}
}

func Exclude(...CommandSurface) Option {

	return func() {}
}

func AllFiles() Files
func AllFunctions() Functions
func Package(name string) PackageInstance

func Evaluate(...Option)
