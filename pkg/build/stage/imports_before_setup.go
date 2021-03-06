package stage

import "github.com/flant/werf/pkg/config"

func GenerateImportsBeforeSetupStage(imageBaseConfig *config.ImageBase, baseStageOptions *NewBaseStageOptions) *ImportsBeforeSetupStage {
	imports := getImports(imageBaseConfig, &getImportsOptions{Before: Setup})
	if len(imports) != 0 {
		return newImportsBeforeSetupStage(imports, baseStageOptions)
	}

	return nil
}

func newImportsBeforeSetupStage(imports []*config.Import, baseStageOptions *NewBaseStageOptions) *ImportsBeforeSetupStage {
	s := &ImportsBeforeSetupStage{}
	s.ImportsStage = newImportsStage(imports, ImportsBeforeSetup, baseStageOptions)
	return s
}

type ImportsBeforeSetupStage struct {
	*ImportsStage
}
