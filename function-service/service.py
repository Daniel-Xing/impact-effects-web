import logging
from concurrent import futures

import grpc
import impactEffect_pb2
import impactEffect_pb2_grpc

import impactEffects.instances.ImpactorClass
from impactEffects.functions import *
from impactEffects.functions.function import *
from impactEffects.instances import ImpactorClass, TargetClass


class ImpactEffectService(impactEffect_pb2_grpc.ImpactEffectServiceServicer):

    def cal_mass(self, request, context):
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)

        return impactEffect_pb2.cal_mass_response(mass=impactor.get_mass())

    def cal_Kinetic_energy(self, request, context):
        print("-------------- cal_Kinetic_energy --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        function_type = Choices.Collins
        kn = kinetic_energy(impactor, function_type)
        print("kinetic_energy: ", kn, "Choice: ", request.choice)
        return impactEffect_pb2.cal_Kinetic_energy_response(Kinetic_energy=kn)

    def cal_kinetic_energy_megatons(self, request, context):
        print("-------------- cal_kinetic_energy_megatons --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        function_type = Choices.Collins
        kn = kinetic_energy_megatons(impactor, function_type)
        print("kinetic_energy: ", kn, "Choice: ", request.choice)
        return impactEffect_pb2.cal_kinetic_energy_megatons_response(kinetic_energy_megatons=kn)

    def cal_rec_time(self, request, context):
        print("-------------- cal_rec_time --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        function_type = Choices.Collins

        print("", impactor.pdiameter, impactor.density,
              impactor.velocity, impactor.theta)
        rec_time_ = rec_time(impactor, function_type)
        print("rec_time: ", rec_time_, "Choice: ", request.choice)

        return impactEffect_pb2.cal_rec_time_response(rec_time=rec_time_)

    def cal_i_factor(self, request, context):
        print("-------------- cal_i_factor --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        collins_iFactor, _av, _rStrength = iFactor(impactor=impactor, target=target,
                                                   type=function_type)

        return impactEffect_pb2.cal_i_factor_response(i_factor=collins_iFactor, av=_av, rStrength=_rStrength)

    def burst_velocity_at_zero(self, request, context):
        print("-------------- burst_velocity_at_zero --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        velocity_at_surface = burst_velocity_at_zero(impactor=impactor, target=target,
                                                     type=function_type)

        return impactEffect_pb2.burst_velocity_at_zero_response(velocity_at_surface=velocity_at_surface)

    def altitude_of_breakup(self, request, context):
        print("-------------- altitude_of_breakup --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        altitudeBU = altitude_of_breakup(impactor=impactor,
                                         target=target,
                                         collins_iFactor=request.collins_iFactor,
                                         type=function_type)

        return impactEffect_pb2.altitude_of_breakup_response(altitudeBU=altitudeBU)

    def velocity_at_breakup(self, request, context):
        print("-------------- velocity_at_breakup --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        velocity = velocity_at_breakup(impactor=impactor,
                                       target=target,
                                       av=request.av,
                                       type=function_type)

        return impactEffect_pb2.velocity_at_breakup_response(velocity=velocity)

    def dispersion_length_scale(self, request, context):
        print("-------------- dispersion_length_scale --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        lDisper = dispersion_length_scale(impactor=impactor,
                                          target=target,
                                          altitudeBU=request.altitudeBU,
                                          type=function_type)

        return impactEffect_pb2.dispersion_length_scale_response(lDisper=lDisper)

    def airburst_altitude(self, request, context):
        print("-------------- airburst_altitude --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        altitudeBurst = airburst_altitude(impactor=impactor,
                                          target=target,
                                          alpha2=request.alpha2,
                                          lDisper=request.lDisper,
                                          type=function_type)

        return impactEffect_pb2.airburst_altitude_response(altitudeBurst=altitudeBurst)

    def brust_velocity(self, request, context):
        print("-------------- brust_velocity --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        velocity = brust_velocity(impactor=impactor,
                                  target=target,
                                  altitudeBurst=request.altitudeBurst,
                                  altitudeBU=request.altitudeBU,
                                  vBu=request.vBu,
                                  lDisper=request.lDisper,
                                  type=function_type)

        return impactEffect_pb2.brust_velocity_response(velocity=velocity)

    def dispersion_of_impactor(self, request, context):
        print("-------------- dispersion_of_impactor --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        dispersion = dispersion_of_impactor(impactor=impactor,
                                            target=target,
                                            l_disper=request.l_disper,
                                            altitude_bu=request.altitude_bu,
                                            altitude_burst=request.altitude_burst,
                                            type=function_type)

        return impactEffect_pb2.dispersion_of_impactor_response(dispersion=dispersion)

    def fraction_of_momentum(self, request, context):
        print("-------------- fraction_of_momentum --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        lratio, pratio = fraction_of_momentum(impactor=impactor,
                                              target=target,
                                              velocity=request.velocity,
                                              type=function_type)

        return impactEffect_pb2.fraction_of_momentum_response(lratio=lratio, pratio=pratio)

    def cal_trot_change(self, request, context):
        print("-------------- cal_trot_change --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        trot_change = cal_trot_change(impactor=impactor,
                                      target=target,
                                      velocity=request.velocity,
                                      type=function_type)

        return impactEffect_pb2.cal_trot_change_response(trot_change=trot_change)

    def cal_energy_atmosphere(self, request, context):
        print("-------------- cal_energy_atmosphere --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        energy_atmosphere = cal_energy_atmosphere(impactor=impactor,
                                                  target=target,
                                                  velocity=request.velocity,
                                                  type=function_type)

        return impactEffect_pb2.cal_energy_atmosphere_response(energy_atmosphere=energy_atmosphere)

    def cal_energy_blast_surface(self, request, context):
        print("-------------- cal_energy_blast_surface --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        energy_blast, energy_surface = cal_energy_blast_surface(impactor=impactor,
                                                                target=target,
                                                                velocity=request.velocity,
                                                                altitudeBurst=request.altitudeBurst,
                                                                energy_atmosphere=request.energy_atmosphere,
                                                                type=function_type)

        return impactEffect_pb2.cal_energy_blast_surface_response(energy_blast=energy_blast, energy_surface=energy_surface)

    def cal_mass_of_water(self, request, context):
        print("-------------- cal_mass_of_water --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        mwater = cal_mass_of_water(impactor=impactor, target=target,
                                   type=function_type)

        return impactEffect_pb2.cal_mass_of_water_response(mwater=mwater)

    def cal_velocity_projectile(self, request, context):
        print("-------------- cal_velocity_projectile --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        vseafloor = cal_velocity_projectile(impactor=impactor, target=target,
                                            type=function_type)

        return impactEffect_pb2.cal_velocity_projectile_response(vseafloor=vseafloor)

    def cal_energy_at_seafloor(self, request, context):
        print("-------------- cal_energy_at_seafloor --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        energy_seafloor = cal_energy_at_seafloor(impactor=impactor, target=target,
                                                 type=function_type)

        return impactEffect_pb2.cal_energy_at_seafloor_response(energy_seafloor=energy_seafloor)

    def cal_ePIcentral_angle(self, request, context):
        print("-------------- cal_ePIcentral_angle --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        delta = cal_ePIcentral_angle(target=target,
                                     type=function_type)

        return impactEffect_pb2.cal_ePIcentral_angle_response(delta=delta)

    def cal_scaling_diameter_constant(self, request, context):
        print("-------------- cal_scaling_diameter_constant --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        Cd, beta = cal_scaling_diameter_constant(target=target,
                                                 type=function_type)

        return impactEffect_pb2.cal_scaling_diameter_constant_response(Cd=Cd, beta=beta)

    def cal_anglefac(self, request, context):
        print("-------------- cal_anglefac --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        anglefac = cal_anglefac(impactor=impactor,
                                type=function_type)

        return impactEffect_pb2.cal_anglefac_response(anglefac=anglefac)

    def cal_wdiameter(self, request, context):
        print("-------------- cal_wdiameter --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        wdiameter = cal_wdiameter(impactor=impactor, target=target,
                                  type=function_type)

        return impactEffect_pb2.cal_wdiameter_response(wdiameter=wdiameter)

    def cal_transient_crater_diameter(self, request, context):
        print("-------------- cal_transient_crater_diameter --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        Dtr = cal_transient_crater_diameter(impactor=impactor, target=target,
                                            type=function_type)

        return impactEffect_pb2.cal_transient_crater_diameter_response(Dtr=Dtr)

    def cal_depthr(self, request, context):
        print("-------------- cal_depthr --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        depthr = cal_depthr(impactor=impactor, target=target,
                            type=function_type)

        return impactEffect_pb2.cal_depthr_response(depthr=depthr)

    def cal_cdiamater(self, request, context):
        print("-------------- cal_cdiamater --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        cdiameter = cal_cdiamater(impactor=impactor, target=target,
                                  type=function_type)

        return impactEffect_pb2.cal_cdiamater_response(cdiameter=cdiameter)

    def cal_brecciaThickness(self, request, context):
        print("-------------- cal_brecciaThickness --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        brecciaThickness = cal_brecciaThickness(impactor=impactor, target=target,
                                                type=function_type)

        return impactEffect_pb2.cal_brecciaThickness_response(brecciaThickness=brecciaThickness)

    def cal_depthfr(self, request, context):
        print("-------------- cal_depthfr --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        depthfr = cal_depthfr(impactor=impactor, target=target,
                              type=function_type)

        return impactEffect_pb2.cal_depthfr_response(depthfr=depthfr)

    def cal_vCrater(self, request, context):
        print("-------------- cal_vCrater --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        vCrater = cal_vCrater(impactor=impactor, target=target,
                              type=function_type)

        return impactEffect_pb2.cal_vCrater_response(vCrater=vCrater)

    def cal_vratio(self, request, context):
        print("-------------- cal_vratio --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        vratio = cal_vratio(impactor=impactor, target=target,
                            type=function_type)

        return impactEffect_pb2.cal_vratio_response(vratio=vratio)

    def cal_vCrater_vRation(self, request, context):
        print("-------------- cal_vCrater_vRation --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        vCrater, vratio = cal_vCrater_vRation(impactor=impactor, target=target,
                                              type=function_type)

        return impactEffect_pb2.cal_vCrater_vRation_response(vCrater=vCrater, vratio=vratio)

    def cal_vMelt(self, request, context):
        print("-------------- cal_vMelt --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        vMelt = cal_vMelt(impactor=impactor, target=target,
                          type=function_type)

        return impactEffect_pb2.cal_vMelt_response(vMelt=vMelt)

    def cal_mratio_and_mcratio(self, request, context):
        print("-------------- cal_mratio_and_mcratio --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        mratio, mcratio = cal_mratio_and_mcratio(impactor=impactor, target=target,
                                                 type=function_type)

        return impactEffect_pb2.cal_mratio_and_mcratio_response(mratio=mratio, mcratio=mcratio)

    def cal_eject_arrival(self, request, context):
        print("-------------- cal_eject_arrival --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        ejecta_arrival = cal_eject_arrival(impactor=impactor, target=target,
                                           type=function_type)

        return impactEffect_pb2.cal_eject_arrival_response(ejecta_arrival=ejecta_arrival)

    def cal_ejecta_thickness(self, request, context):
        print("-------------- cal_ejecta_thickness --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        ejecta_thickness = cal_ejecta_thickness(impactor=impactor, target=target,
                                                type=function_type)

        return impactEffect_pb2.cal_ejecta_thickness_response(ejecta_thickness=ejecta_thickness)

    def cal_d_frag(self, request, context):
        print("-------------- cal_d_frag --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        d_frag = cal_d_frag(impactor=impactor, target=target,
                            type=function_type)

        return impactEffect_pb2.cal_d_frag_response(d_frag=d_frag)

    def cal_themal(self, request, context):
        print("-------------- cal_themal --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        (
            h,
            Rf,
            thermal_exposure_,
            no_radiation_,
            max_rad_time_,
            irradiation_time_,
            megaton_factor_,
            thermal_power_,
        ) = cal_themal(impactor=impactor, target=target,
                       type=function_type)

        return impactEffect_pb2.cal_themal_response(h=h,
                                                    Rf=Rf,
                                                    thermal_exposure=thermal_exposure_,
                                                    no_radiation=no_radiation_,
                                                    max_rad_time=max_rad_time_,
                                                    irradiation_time=irradiation_time_,
                                                    megaton_factor=megaton_factor_,
                                                    thermal_power=thermal_power_)

    def cal_magnitude(self, request, context):
        print("-------------- cal_magnitude --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        magnitude = cal_magnitude(impactor=impactor, target=target,
                                  type=function_type)

        return impactEffect_pb2.cal_magnitude_response(magnitude=magnitude)

    def cal_magnitude2(self, request, context):
        print("-------------- cal_magnitude2 --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        eff_mag, seismic_arrival = cal_magnitude2(impactor=impactor, target=target,
                                                  type=function_type)

        return impactEffect_pb2.cal_magnitude2_response(eff_mag=eff_mag, seismic_arrival=seismic_arrival)

    def cal_shock_arrival(self, request, context):
        print("-------------- cal_shock_arrival --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        shock_arrival = cal_shock_arrival(impactor=impactor, target=target,
                                          type=function_type)

        return impactEffect_pb2.cal_shock_arrival_response(shock_arrival=shock_arrival)

    def cal_vmax(self, request, context):
        print("-------------- cal_vmax --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        vmax, opressure = cal_vmax(impactor=impactor, target=target,
                                   type=function_type)

        return impactEffect_pb2.cal_vmax_response(vmax=vmax, opressure=opressure)

    def cal_shock_damage(self, request, context):
        print("-------------- cal_shock_damage --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        shock_damage = cal_shock_damage(impactor=impactor, target=target,
                                        type=function_type)

        return impactEffect_pb2.cal_shock_damage_response(shock_damage=shock_damage)

    def cal_dec_level(self, request, context):
        print("-------------- cal_dec_level --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        dec_level = cal_dec_level(impactor=impactor, target=target,
                                  type=function_type)

        return impactEffect_pb2.cal_dec_level_response(dec_level=dec_level)

    def cal_TsunamiArrivalTime(self, request, context):
        print("-------------- cal_TsunamiArrivalTime --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        TsunamiArrivalTime = cal_TsunamiArrivalTime(impactor=impactor, target=target,
                                                    type=function_type)

        return impactEffect_pb2.cal_TsunamiArrivalTime_response(TsunamiArrivalTime=TsunamiArrivalTime)

    def cal_WaveAmplitudeUpperLimit(self, request, context):
        print("-------------- cal_WaveAmplitudeUpperLimit --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        WaveAmplitudeUpperLimit = cal_WaveAmplitudeUpperLimit(impactor=impactor, target=target,
                                                              type=function_type)

        return impactEffect_pb2.cal_WaveAmplitudeUpperLimit_response(WaveAmplitudeUpperLimit=WaveAmplitudeUpperLimit)

    def cal_WaveAmplitudeLowerLimit(self, request, context):
        print("-------------- cal_WaveAmplitudeLowerLimit --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.impactor.diameter,
                            density=request.impactor.density,
                            velocity=request.impactor.velocity,
                            theta=request.impactor.theta)
        target = Target(depth=request.targets.depth,
                        distance=request.targets.distance,
                        density=request.targets.density)
        function_type = Choices.Collins

        WaveAmplitudeLowerLimit = cal_WaveAmplitudeLowerLimit(impactor=impactor, target=target,
                                                              type=function_type)

        return impactEffect_pb2.cal_WaveAmplitudeLowerLimit_response(WaveAmplitudeLowerLimit=WaveAmplitudeLowerLimit)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    impactEffect_pb2_grpc.add_ImpactEffectServiceServicer_to_server(
        ImpactEffectService(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()
