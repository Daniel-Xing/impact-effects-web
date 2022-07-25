package controlor

import (
	"back-web/google.golang.org/grpc/impactEffect/impactEffect"
	"back-web/rpc"
	"log"
	"math"
)

func packImapctEffectArgs() (*impactEffect.Impactor, *impactEffect.Targets) {
	impactor := &impactEffect.Impactor{}
	impactor.Density = 111000
	impactor.Diameter = 111
	impactor.Velocity = 111
	impactor.Theta = 45

	target := &impactEffect.Targets{}
	target.Density = 111
	target.Depth = 0
	target.Distance = 111

	return impactor, target
}

func ImpactEffect() {
	impactor, target := packImapctEffectArgs()

	// calculate the ennergy
	ies := rpc.NewImpactEffectRPCService()
	defer ies.Close()
	// cal_energy
	_kinetic_energy := ies.Cal_KineticEnergy(&impactEffect.Cal_KineticEnergyRequest{
		Impactor: impactor,
		Choice:   1,
	})
	log.Println("_kinetic_energy", _kinetic_energy)

	// cal Kinetic Energy Megatons
	_kinetic_energy_megatons := ies.CalKineticEnergyMegatons(&impactEffect.CalKineticEnergyMegatonsRequest{
		Impactor: impactor,
		Choice:   1,
	})
	log.Println("_kinetic_energy_megatons: ", _kinetic_energy_megatons)

	// calculate rec time
	_rec_time := ies.CalRecTime(&impactEffect.CalRecTimeRequest{
		Impactor: impactor,
		Choice:   1,
	})
	log.Println("_rec_time:", _rec_time)

	// calculate i Factor
	collins_iFactor, av, rStrength := ies.CalIFactor(&impactEffect.CalIFactorRequest{
		Impactor: impactor,
		Targets:  target,
		Choice:   1,
	})
	log.Println("collins_iFactor:", collins_iFactor)

	//
	var velocity, altitudeBU, vBU, lDisper, altitudeBurst, dispersion float32
	if collins_iFactor >= 1 {
		velocity = ies.BurstVelocityAtZero(&impactEffect.BurstVelocityAtZeroRequest{
			Impactor: impactor,
			Targets:  target,
			Choice:   1,
		})
	} else {
		altitudeBU = ies.AltitudeOfBreakup(&impactEffect.AltitudeOfBreakupRequest{
			Impactor:  impactor,
			Targets:   target,
			Choice:    1,
			RStrength: rStrength,
		})

		vBU = ies.VelocityAtBreakup(&impactEffect.VelocityAtBreakupRequest{
			Impactor:   impactor,
			Targets:    target,
			Choice:     1,
			Av:         av,
			AltitudeBU: altitudeBU,
		})

		lDisper = ies.DispersionLengthScale(&impactEffect.DispersionLengthScaleRequest{
			Impactor:   impactor,
			Targets:    target,
			Choice:     1,
			AltitudeBU: altitudeBU,
		})

		altitudeBurst = ies.AirburstAltitude(&impactEffect.AirburstAltitudeRequest{
			Impactor: impactor,
			Targets:  target,
			Choice:   1,
			LDisper:  lDisper,
		})

		velocity = ies.BrustVelocity(&impactEffect.BrustVelocityRequest{
			Impactor:      impactor,
			Targets:       target,
			Choice:        1,
			AltitudeBurst: altitudeBurst,
			AltitudeBU:    altitudeBU,
			VBu:           vBU,
			LDisper:       lDisper,
		})

		dispersion = ies.DispersionOfImpactor(&impactEffect.DispersionOfImpactorRequest{
			Impactor:      impactor,
			Targets:       target,
			Choice:        1,
			LDisper:       lDisper,
			AltitudeBu:    altitudeBU,
			AltitudeBurst: altitudeBurst,
		})
	}

	log.Println("velocity, altitudeBU, vBU, lDisper, altitudeBurst, dispersion: ", velocity, altitudeBU, vBU, lDisper, altitudeBurst, dispersion)

	lratio, pratio := ies.FractionOfMomentum(&impactEffect.FractionOfMomentumRequest{
		Impactor: impactor,
		Targets:  target,
		Choice:   1,
		Velocity: velocity,
	})
	log.Println("lratio, pratio: ", lratio, pratio)

	trot_change := ies.CalTrotChange(&impactEffect.CalTrotChangeRequest{
		Impactor: impactor,
		Targets:  target,
		Choice:   1,
		Velocity: velocity,
	})
	log.Println("trot_change", trot_change)

	energy_atmosphere := ies.CalEnergyAtmosphere(&impactEffect.CalEnergyAtmosphereRequest{
		Impactor: impactor,
		Targets:  target,
		Choice:   1,
		Velocity: velocity,
	})
	log.Println("energy_atmosphere", energy_atmosphere)

	energy_blast, energy_surface := ies.CalEnergyBlastSurface(&impactEffect.CalEnergyBlastSurfaceRequest{
		Impactor:         impactor,
		Targets:          target,
		Choice:           1,
		Velocity:         velocity,
		AltitudeBurst:    altitudeBurst,
		EnergyAtmosphere: energy_atmosphere,
	})
	print(energy_blast, energy_surface)

	mwater := ies.CalMassOfWater(&impactEffect.CalMassOfWaterRequest{
		Impactor: impactor,
		Targets:  target,
		Choice:   1,
	})
	print("mwater: ", mwater)

	vseafloor := ies.CalVelocityProjectile(&impactEffect.CalVelocityProjectileRequest{
		Impactor: impactor,
		Targets:  target,
		Choice:   1,
		Velocity: velocity,
	})
	print("vseafloor: ", vseafloor)

	energy_seafloor := ies.CalEnergyAtSeafloor(&impactEffect.CalEnergyAtSeafloorRequest{
		Impactor:  impactor,
		Targets:   target,
		Choice:    1,
		Vseafloor: vseafloor,
	})
	print("energy_seafloor:", energy_seafloor)

	delta := ies.CalEPIcentralAngle(&impactEffect.CalEPIcentralAngleRequest{
		Impactor: impactor,
		Targets:  target,
		Choice:   1,
	})
	print("delta: ", delta)

	cd, beta := ies.CalScalingDiameterConstant(&impactEffect.CalScalingDiameterConstantRequest{
		Impactor: impactor,
		Targets:  target,
		Choice:   1,
	})
	print("cd, beta: ", cd, beta)

	anglefac := ies.CalAnglefac(&impactEffect.CalAnglefacRequest{
		Impactor: impactor,
		Targets:  target,
		Choice:   1,
	})
	print("anglefac: ", anglefac)

	var wdiameter float32 = 0
	if target.Depth != 0 {
		wdiameter = ies.CalWdiameter(&impactEffect.CalWdiameterRequest{
			Impactor: impactor,
			Targets:  target,
			Choice:   1,
			Anglefac: anglefac,
			Velocity: velocity,
		})
	}
	log.Println("wdiameter: ", wdiameter)

	Dtr := ies.CalTransientCraterDiameter(&impactEffect.CalTransientCraterDiameterRequest{
		Impactor:   impactor,
		Targets:    target,
		Choice:     1,
		Cd:         cd,
		Beta:       beta,
		Anglefac:   anglefac,
		Vseafloor:  vseafloor,
		Dispersion: dispersion,
	})
	log.Println("Dtr: ", Dtr)

	depthr := ies.CalDepthr(&impactEffect.CalDepthrRequest{
		Impactor: impactor,
		Targets:  target,
		Choice:   1,
		Dtr:      Dtr,
	})
	log.Println("depthr: ", depthr)

	vCrater, vratio := ies.CalVCraterVRation(&impactEffect.CalVCraterVRationRequest{
		Impactor: impactor,
		Targets:  target,
		Choice:   1,
		Dtr:      Dtr,
	})
	log.Println("vCrater, vratio: ", vCrater, vratio)

	vMelt := ies.CalVMelt(&impactEffect.CalVMeltRequest{
		Impactor:       impactor,
		Targets:        target,
		Choice:         1,
		Velocity:       velocity,
		EnergySeafloor: energy_seafloor,
	})
	log.Println("vMelt: ", vMelt)

	mratio, mcratio := ies.CalMratioAndMcratio(&impactEffect.CalMratioAndMcratioRequest{
		Impactor: impactor,
		Targets:  target,
		Choice:   1,
		Velocity: velocity,
		VMelt:    vMelt,
		VCrater:  vCrater,
		Dtr:      Dtr,
	})
	log.Println("mratio, mcratio: ", mratio, mcratio)

	energy_megatons := energy_surface / float32(4.186*math.Pow(10, 15))

	mass := ies.Cal_mass(&impactEffect.CalMassRequest{
		Impactor: impactor,
	})

	if mass <= 1.5707963e12 {
		log.Println(mass, impactor.Velocity, velocity, collins_iFactor, altitudeBU,
			altitudeBurst, impactor.Density, dispersion, impactor.Theta, energy_surface, energy_megatons)
	}

	if altitudeBurst <= 0 {
		ejecta_arrival := ies.CalEjectArrival(&impactEffect.CalEjectArrivalRequest{
			Impactor:      impactor,
			Targets:       target,
			Choice:        1,
			AltitudeBurst: altitudeBurst,
		})
		log.Println("ejecta_arrival: ", ejecta_arrival)

		ejecta_thickness := ies.CalEjectaThickness(&impactEffect.CalEjectaThicknessRequest{
			Impactor:      impactor,
			Targets:       target,
			Choice:        1,
			AltitudeBurst: altitudeBurst,
			Dtr:           Dtr,
		})
		log.Println("ejecta_thickness: ", ejecta_thickness)

		d_frag := ies.CalDFrag(&impactEffect.CalDFragRequest{
			Impactor:      impactor,
			Targets:       target,
			Choice:        1,
			AltitudeBurst: altitudeBurst,
			Dtr:           Dtr,
		})
		log.Println("d_frag: ", d_frag)

		if velocity >= 15 {
			h, Rf, thermal_exposure_, no_radiation_, max_rad_time_, irradiation_time_, megaton_factor_, thermal_power_ := ies.CalThemal(&impactEffect.CalThemalRequest{
				Impactor:      impactor,
				Targets:       target,
				Choice:        1,
				AltitudeBurst: altitudeBurst,
				Velocity:      velocity,
				Delta:         delta,
				EnergySurface: energy_surface,
			})
			print("h, Rf, thermal_exposure_, no_radiation_, max_rad_time_, irradiation_time_, megaton_factor_, thermal_power_ : ",
				h, Rf, thermal_exposure_, no_radiation_, max_rad_time_, irradiation_time_, megaton_factor_, thermal_power_)
		}

		magnitude := ies.CalMagnitude(&impactEffect.CalMagnitudeRequest{
			Impactor:       impactor,
			Targets:        target,
			Choice:         1,
			AltitudeBurst:  altitudeBurst,
			EnergySeafloor: energy_seafloor,
		})
		log.Println("magnitude: ", magnitude)

		eff_mag, seismic_arrival := ies.CalMagnitude2(&impactEffect.CalMagnitude2Request{
			Impactor:       impactor,
			Targets:        target,
			Choice:         1,
			AltitudeBurst:  altitudeBurst,
			EnergySeafloor: energy_seafloor,
			Delta:          delta,
		})
		log.Println("eff_mag, seismic_arrival: ", eff_mag, seismic_arrival)

		if target.Distance*1000 <= Dtr/2 {
			log.Println("Exit")
			return
		}
	}

	shock_arrival := ies.CalShockArrival(&impactEffect.CalShockArrivalRequest{
		Impactor:      impactor,
		Targets:       target,
		Choice:        1,
		AltitudeBurst: altitudeBurst,
	})
	log.Println("shock_arrival: ", shock_arrival)

	vmax, opressure := ies.CalVmax(&impactEffect.CalVmaxRequest{
		Impactor:      impactor,
		Targets:       target,
		Choice:        1,
		AltitudeBurst: altitudeBurst,
		EnergyBlast:   energy_blast,
	})

	shock_damage := ies.CalShockDamage(&impactEffect.CalShockDamageRequest{
		Impactor:  impactor,
		Targets:   target,
		Choice:    1,
		Opressure: opressure,
		Vmax:      vmax,
	})
	log.Println("shock_damage: ", shock_damage)

	dec_level := ies.CalDecLevel(&impactEffect.CalDecLevelRequest{
		Impactor:      impactor,
		Targets:       target,
		Choice:        1,
		AltitudeBurst: altitudeBurst,
		EnergyBlast:   energy_blast,
	})
	log.Println("dec_level: ", dec_level)

	if target.Depth > 0 {
		TsunamiArrivalTime := ies.Cal_TsunamiArrivalTime(&impactEffect.Cal_TsunamiArrivalTimeRequest{
			Impactor:  impactor,
			Targets:   target,
			Choice:    1,
			Wdiameter: wdiameter,
		})
		log.Println("TsunamiArrivalTime: ", TsunamiArrivalTime)

		WaveAmplitudeUpperLimit := ies.Cal_WaveAmplitudeUpperLimit(&impactEffect.Cal_WaveAmplitudeUpperLimitRequest{
			Impactor:  impactor,
			Targets:   target,
			Choice:    1,
			Wdiameter: wdiameter,
		})
		log.Println("WaveAmplitudeUpperLimit: ", WaveAmplitudeUpperLimit)

		WaveAmplitudeLowerLimit := ies.Cal_WaveAmplitudeLowerLimit(&impactEffect.Cal_WaveAmplitudeLowerLimitRequest{
			Impactor:  impactor,
			Targets:   target,
			Choice:    1,
			Wdiameter: wdiameter,
		})
		log.Println("WaveAmplitudeLowerLimit: ", WaveAmplitudeLowerLimit)
	}
}
