package controlor

import (
	"back-web/cache"
	"back-web/google.golang.org/grpc/impactEffect/impactEffect"
	"back-web/model"
	"back-web/rpc"
	"back-web/util"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/gin-gonic/gin"
)

func packImapctEffectArgs(ctx *gin.Context) (*impactEffect.Impactor, *impactEffect.Targets) {

	var requestMap = model.Impact{}
	json.NewDecoder(ctx.Request.Body).Decode(&requestMap)

	impactor := &impactEffect.Impactor{}
	impactor.Density = requestMap.ImpactorDensity
	impactor.Diameter = requestMap.ImpactorDiameter
	impactor.Velocity = requestMap.ImpactorVelocity
	impactor.Theta = requestMap.ImpactorTheta

	target := &impactEffect.Targets{}
	target.Density = requestMap.TargetDensity
	target.Depth = requestMap.TargetDepth
	target.Distance = requestMap.TargetDistance

	return impactor, target
}

func ImpactEffect(ctx *gin.Context) {
	defer util.Success(ctx, "okokokokokokokoko", "SUCCESS")
	// log
	log.Println("get the request")
	impactor, target := packImapctEffectArgs(ctx)

	// read from cache
	redisClient := cache.GetCache()
	RedisUtilInstance := cache.RedisUtilInstance(redisClient)
	result, err := RedisUtilInstance.HGet("imapctEffect", fmt.Sprintf("%f_%f_%f_%f_%f_%f_%f",
		impactor.Density, impactor.Diameter, impactor.Velocity, impactor.Theta, target.Density, target.Depth, target.Density))
	if err == nil && result != "" {
		log.Println("Read from Redis")
		return
	}

	// calculate the ennergy
	ies, err := rpc.NewImpactEffectRPCService()
	if err != nil {
		log.Println(err)
		return
	}
	defer ies.Close()

	// cal_energy
	_kinetic_energy, err := ies.Cal_KineticEnergy(&impactEffect.Cal_KineticEnergyRequest{
		Impactor: impactor,
		Choice:   1,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("_kinetic_energy", _kinetic_energy)

	// cal Kinetic Energy Megatons
	_kinetic_energy_megatons, err := ies.CalKineticEnergyMegatons(&impactEffect.CalKineticEnergyMegatonsRequest{
		Impactor: impactor,
		Choice:   1,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("_kinetic_energy_megatons: ", _kinetic_energy_megatons)

	// calculate rec time
	_rec_time, err := ies.CalRecTime(&impactEffect.CalRecTimeRequest{
		Impactor: impactor,
		Choice:   1,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("_rec_time:", _rec_time)

	// calculate i Factor
	collins_iFactor, av, rStrength, err := ies.CalIFactor(&impactEffect.CalIFactorRequest{
		Impactor: impactor,
		Targets:  target,
		Choice:   1,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("collins_iFactor:", collins_iFactor)

	//
	var velocity, altitudeBU, vBU, lDisper, altitudeBurst, dispersion float32
	if collins_iFactor >= 1 {
		velocity, err = ies.BurstVelocityAtZero(&impactEffect.BurstVelocityAtZeroRequest{
			Impactor: impactor,
			Targets:  target,
			Choice:   1,
		})
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		altitudeBU, err = ies.AltitudeOfBreakup(&impactEffect.AltitudeOfBreakupRequest{
			Impactor:  impactor,
			Targets:   target,
			Choice:    1,
			RStrength: rStrength,
		})
		if err != nil {
			log.Println(err)
			return
		}

		vBU, err = ies.VelocityAtBreakup(&impactEffect.VelocityAtBreakupRequest{
			Impactor:   impactor,
			Targets:    target,
			Choice:     1,
			Av:         av,
			AltitudeBU: altitudeBU,
		})
		if err != nil {
			log.Println(err)
			return
		}

		lDisper, err = ies.DispersionLengthScale(&impactEffect.DispersionLengthScaleRequest{
			Impactor:   impactor,
			Targets:    target,
			Choice:     1,
			AltitudeBU: altitudeBU,
		})
		if err != nil {
			log.Println(err)
			return
		}

		altitudeBurst, err = ies.AirburstAltitude(&impactEffect.AirburstAltitudeRequest{
			Impactor: impactor,
			Targets:  target,
			Choice:   1,
			LDisper:  lDisper,
		})
		if err != nil {
			log.Println(err)
			return
		}

		velocity, err = ies.BrustVelocity(&impactEffect.BrustVelocityRequest{
			Impactor:      impactor,
			Targets:       target,
			Choice:        1,
			AltitudeBurst: altitudeBurst,
			AltitudeBU:    altitudeBU,
			VBu:           vBU,
			LDisper:       lDisper,
		})
		if err != nil {
			log.Println(err)
			return
		}

		dispersion, err = ies.DispersionOfImpactor(&impactEffect.DispersionOfImpactorRequest{
			Impactor:      impactor,
			Targets:       target,
			Choice:        1,
			LDisper:       lDisper,
			AltitudeBu:    altitudeBU,
			AltitudeBurst: altitudeBurst,
		})
		if err != nil {
			log.Println(err)
			return
		}
	}

	log.Println("velocity, altitudeBU, vBU, lDisper, altitudeBurst, dispersion: ", velocity, altitudeBU, vBU, lDisper, altitudeBurst, dispersion)

	lratio, pratio, err := ies.FractionOfMomentum(&impactEffect.FractionOfMomentumRequest{
		Impactor: impactor,
		Targets:  target,
		Choice:   1,
		Velocity: velocity,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("lratio, pratio: ", lratio, pratio)

	trot_change, err := ies.CalTrotChange(&impactEffect.CalTrotChangeRequest{
		Impactor: impactor,
		Targets:  target,
		Choice:   1,
		Velocity: velocity,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("trot_change", trot_change)

	energy_atmosphere, err := ies.CalEnergyAtmosphere(&impactEffect.CalEnergyAtmosphereRequest{
		Impactor: impactor,
		Targets:  target,
		Choice:   1,
		Velocity: velocity,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("energy_atmosphere", energy_atmosphere)

	energy_blast, energy_surface, err := ies.CalEnergyBlastSurface(&impactEffect.CalEnergyBlastSurfaceRequest{
		Impactor:         impactor,
		Targets:          target,
		Choice:           1,
		Velocity:         velocity,
		AltitudeBurst:    altitudeBurst,
		EnergyAtmosphere: energy_atmosphere,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(energy_blast, energy_surface)

	mwater, err := ies.CalMassOfWater(&impactEffect.CalMassOfWaterRequest{
		Impactor: impactor,
		Targets:  target,
		Choice:   1,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("mwater: ", mwater)

	vseafloor, err := ies.CalVelocityProjectile(&impactEffect.CalVelocityProjectileRequest{
		Impactor: impactor,
		Targets:  target,
		Choice:   1,
		Velocity: velocity,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("vseafloor: ", vseafloor)

	energy_seafloor, err := ies.CalEnergyAtSeafloor(&impactEffect.CalEnergyAtSeafloorRequest{
		Impactor:  impactor,
		Targets:   target,
		Choice:    1,
		Vseafloor: vseafloor,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("energy_seafloor:", energy_seafloor)

	delta, err := ies.CalEPIcentralAngle(&impactEffect.CalEPIcentralAngleRequest{
		Impactor: impactor,
		Targets:  target,
		Choice:   1,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("delta: ", delta)

	cd, beta, err := ies.CalScalingDiameterConstant(&impactEffect.CalScalingDiameterConstantRequest{
		Impactor: impactor,
		Targets:  target,
		Choice:   1,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("cd, beta: ", cd, beta)

	anglefac, err := ies.CalAnglefac(&impactEffect.CalAnglefacRequest{
		Impactor: impactor,
		Targets:  target,
		Choice:   1,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("anglefac: ", anglefac)

	var wdiameter float32 = 0
	if target.Depth != 0 {
		wdiameter, err = ies.CalWdiameter(&impactEffect.CalWdiameterRequest{
			Impactor: impactor,
			Targets:  target,
			Choice:   1,
			Anglefac: anglefac,
			Velocity: velocity,
		})
	}
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("wdiameter: ", wdiameter)

	Dtr, err := ies.CalTransientCraterDiameter(&impactEffect.CalTransientCraterDiameterRequest{
		Impactor:   impactor,
		Targets:    target,
		Choice:     1,
		Cd:         cd,
		Beta:       beta,
		Anglefac:   anglefac,
		Vseafloor:  vseafloor,
		Dispersion: dispersion,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Dtr: ", Dtr)

	depthr, err := ies.CalDepthr(&impactEffect.CalDepthrRequest{
		Impactor: impactor,
		Targets:  target,
		Choice:   1,
		Dtr:      Dtr,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("depthr: ", depthr)

	vCrater, vratio, err := ies.CalVCraterVRation(&impactEffect.CalVCraterVRationRequest{
		Impactor: impactor,
		Targets:  target,
		Choice:   1,
		Dtr:      Dtr,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("vCrater, vratio: ", vCrater, vratio)

	vMelt, err := ies.CalVMelt(&impactEffect.CalVMeltRequest{
		Impactor:       impactor,
		Targets:        target,
		Choice:         1,
		Velocity:       velocity,
		EnergySeafloor: energy_seafloor,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("vMelt: ", vMelt)

	mratio, mcratio, err := ies.CalMratioAndMcratio(&impactEffect.CalMratioAndMcratioRequest{
		Impactor: impactor,
		Targets:  target,
		Choice:   1,
		Velocity: velocity,
		VMelt:    vMelt,
		VCrater:  vCrater,
		Dtr:      Dtr,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("mratio, mcratio: ", mratio, mcratio)

	energy_megatons := energy_surface / float32(4.186*math.Pow(10, 15))

	mass, err := ies.Cal_mass(&impactEffect.CalMassRequest{
		Impactor: impactor,
	})
	if err != nil {
		log.Println(err)
		return
	}

	if mass <= 1.5707963e12 {
		log.Println(mass, impactor.Velocity, velocity, collins_iFactor, altitudeBU,
			altitudeBurst, impactor.Density, dispersion, impactor.Theta, energy_surface, energy_megatons)
	}

	if altitudeBurst <= 0 {
		ejecta_arrival, err := ies.CalEjectArrival(&impactEffect.CalEjectArrivalRequest{
			Impactor:      impactor,
			Targets:       target,
			Choice:        1,
			AltitudeBurst: altitudeBurst,
		})
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("ejecta_arrival: ", ejecta_arrival)

		ejecta_thickness, err := ies.CalEjectaThickness(&impactEffect.CalEjectaThicknessRequest{
			Impactor:      impactor,
			Targets:       target,
			Choice:        1,
			AltitudeBurst: altitudeBurst,
			Dtr:           Dtr,
		})
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("ejecta_thickness: ", ejecta_thickness)

		d_frag, err := ies.CalDFrag(&impactEffect.CalDFragRequest{
			Impactor:      impactor,
			Targets:       target,
			Choice:        1,
			AltitudeBurst: altitudeBurst,
			Dtr:           Dtr,
		})
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("d_frag: ", d_frag)

		if velocity >= 15 {
			h, Rf, thermal_exposure_, no_radiation_, max_rad_time_, irradiation_time_, megaton_factor_, thermal_power_, err := ies.CalThemal(&impactEffect.CalThemalRequest{
				Impactor:      impactor,
				Targets:       target,
				Choice:        1,
				AltitudeBurst: altitudeBurst,
				Velocity:      velocity,
				Delta:         delta,
				EnergySurface: energy_surface,
			})
			if err != nil {
				log.Println(err)
				return
			}
			log.Println("h, Rf, thermal_exposure_, no_radiation_, max_rad_time_, irradiation_time_, megaton_factor_, thermal_power_ : ",
				h, Rf, thermal_exposure_, no_radiation_, max_rad_time_, irradiation_time_, megaton_factor_, thermal_power_)
		}

		magnitude, err := ies.CalMagnitude(&impactEffect.CalMagnitudeRequest{
			Impactor:       impactor,
			Targets:        target,
			Choice:         1,
			AltitudeBurst:  altitudeBurst,
			EnergySeafloor: energy_seafloor,
		})
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("magnitude: ", magnitude)

		eff_mag, seismic_arrival, err := ies.CalMagnitude2(&impactEffect.CalMagnitude2Request{
			Impactor:       impactor,
			Targets:        target,
			Choice:         1,
			AltitudeBurst:  altitudeBurst,
			EnergySeafloor: energy_seafloor,
			Delta:          delta,
		})
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("eff_mag, seismic_arrival: ", eff_mag, seismic_arrival)

		if target.Distance*1000 <= Dtr/2 {
			log.Println("Exit")
			return
		}
	}

	shock_arrival, err := ies.CalShockArrival(&impactEffect.CalShockArrivalRequest{
		Impactor:      impactor,
		Targets:       target,
		Choice:        1,
		AltitudeBurst: altitudeBurst,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("shock_arrival: ", shock_arrival)

	vmax, opressure, err := ies.CalVmax(&impactEffect.CalVmaxRequest{
		Impactor:      impactor,
		Targets:       target,
		Choice:        1,
		AltitudeBurst: altitudeBurst,
		EnergyBlast:   energy_blast,
	})
	if err != nil {
		log.Println(err)
		return
	}

	shock_damage, err := ies.CalShockDamage(&impactEffect.CalShockDamageRequest{
		Impactor:  impactor,
		Targets:   target,
		Choice:    1,
		Opressure: opressure,
		Vmax:      vmax,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("shock_damage: ", shock_damage)

	dec_level, err := ies.CalDecLevel(&impactEffect.CalDecLevelRequest{
		Impactor:      impactor,
		Targets:       target,
		Choice:        1,
		AltitudeBurst: altitudeBurst,
		EnergyBlast:   energy_blast,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("dec_level: ", dec_level)

	if target.Depth > 0 {
		TsunamiArrivalTime, err := ies.Cal_TsunamiArrivalTime(&impactEffect.Cal_TsunamiArrivalTimeRequest{
			Impactor:  impactor,
			Targets:   target,
			Choice:    1,
			Wdiameter: wdiameter,
		})
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("TsunamiArrivalTime: ", TsunamiArrivalTime)

		WaveAmplitudeUpperLimit, err := ies.Cal_WaveAmplitudeUpperLimit(&impactEffect.Cal_WaveAmplitudeUpperLimitRequest{
			Impactor:  impactor,
			Targets:   target,
			Choice:    1,
			Wdiameter: wdiameter,
		})
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("WaveAmplitudeUpperLimit: ", WaveAmplitudeUpperLimit)

		WaveAmplitudeLowerLimit, err := ies.Cal_WaveAmplitudeLowerLimit(&impactEffect.Cal_WaveAmplitudeLowerLimitRequest{
			Impactor:  impactor,
			Targets:   target,
			Choice:    1,
			Wdiameter: wdiameter,
		})
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("WaveAmplitudeLowerLimit: ", WaveAmplitudeLowerLimit)
	}

	err = RedisUtilInstance.HSetWithExpirationTime("imapctEffect", "111000, 111, 111, 45, 111, 0, 111", "I am here", 60*time.Minute)
	if err != nil {
		log.Println("Set Redis Error")
		return
	}

}
