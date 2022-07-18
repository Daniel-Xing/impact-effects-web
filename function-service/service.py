from concurrent import futures
import imp
import logging
import math
import time

import grpc
import impactEffect_pb2
import impactEffect_pb2_grpc

import impactEffects.instances.ImpactorClass
from impactEffects.functions import *
from impactEffects.functions.function import *
from impactEffects.instances import ImpactorClass, TargetClass


class ImpactEffectService(impactEffect_pb2_grpc.ImpactEffectServiceServicer):

    def cal_Kinetic_energy(self, request, context):
        print("-------------- GetKineticEnergy --------------")
        # impactor = request.impactor
        impactor = Impactor(diameter=request.diameter, density=request.density, velocity=request.velocity,
                            theta=request.theta, depth=request.depth, ttype=request.ttype)
        kn = kinetic_energy(impactor)
        return impactEffect_pb2.KineticEnergy(kinetic_energy=kn)


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
