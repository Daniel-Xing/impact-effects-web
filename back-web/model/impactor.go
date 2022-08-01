package model

type Impact struct {
	ImpactorDensity  float32 `from:"imapctor_density" json:"impactor_density"`
	ImpactorDiameter float32 `from:"imapctor_diameter" json:"impactor_diameter"`
	ImpactorVelocity float32 `from:"imapctor_velocity" json:"impactor_velocity"`
	ImpactorTheta    float32 `from:"imapctor_theta" json:"impactor_theta"`
	TargetDensity    float32 `from:"target_density" json:"target_density"`
	TargetDepth      float32 `from:"target_depth" json:"target_depth"`
	TargetDistance   float32 `from:"target_distance" json:"target_distance"`
}
