package rpc

import (
	"back-web/google.golang.org/grpc/impactEffect/impactEffect"
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/examples/data"
)

type ImpactEffectRPCService struct {
	client impactEffect.ImpactEffectServiceClient
	conn   *grpc.ClientConn
}

/*
Config: TODO： choose a better way
*/
var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func NewImpactEffectRPCService() *ImpactEffectRPCService {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = data.Path("x509/ca_cert.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	// defer conn.Close()
	client := impactEffect.NewImpactEffectServiceClient(conn)

	return &ImpactEffectRPCService{client: client, conn: conn}
}

func (ies *ImpactEffectRPCService) Close() error {
	return ies.conn.Close()
}

func (ies *ImpactEffectRPCService) Cal_mass(req *impactEffect.CalMassRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalMass(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetMass()
}

func (ies *ImpactEffectRPCService) Cal_KineticEnergy(req *impactEffect.Cal_KineticEnergyRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.Cal_KineticEnergy(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetKineticEnergy()
}

func (ies *ImpactEffectRPCService) CalKineticEnergyMegatons(req *impactEffect.CalKineticEnergyMegatonsRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalKineticEnergyMegatons(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetKineticEnergyMegatons()
}

func (ies *ImpactEffectRPCService) CalRecTime(req *impactEffect.CalRecTimeRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalRecTime(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetRecTime()
}

func (ies *ImpactEffectRPCService) CalIFactor(req *impactEffect.CalIFactorRequest) (float32, float32, float32) {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalIFactor(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetIFactor(), result.GetAv(), result.GetRStrength()
}

func (ies *ImpactEffectRPCService) BurstVelocityAtZero(req *impactEffect.BurstVelocityAtZeroRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.BurstVelocityAtZero(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetVelocityAtSurface()
}

func (ies *ImpactEffectRPCService) AltitudeOfBreakup(req *impactEffect.AltitudeOfBreakupRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.AltitudeOfBreakup(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetAltitudeBU()
}

func (ies *ImpactEffectRPCService) VelocityAtBreakup(req *impactEffect.VelocityAtBreakupRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.VelocityAtBreakup(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetVelocity()
}

func (ies *ImpactEffectRPCService) DispersionLengthScale(req *impactEffect.DispersionLengthScaleRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.DispersionLengthScale(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetLDisper()
}

func (ies *ImpactEffectRPCService) AirburstAltitude(req *impactEffect.AirburstAltitudeRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.AirburstAltitude(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetAltitudeBurst()
}

func (ies *ImpactEffectRPCService) BrustVelocity(req *impactEffect.BrustVelocityRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.BrustVelocity(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetVelocity()
}

func (ies *ImpactEffectRPCService) DispersionOfImpactor(req *impactEffect.DispersionOfImpactorRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.DispersionOfImpactor(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetDispersion()
}

func (ies *ImpactEffectRPCService) FractionOfMomentum(req *impactEffect.FractionOfMomentumRequest) (float32, float32) {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.FractionOfMomentum(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetLratio(), result.GetPratio()
}

func (ies *ImpactEffectRPCService) CalTrotChange(req *impactEffect.CalTrotChangeRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalTrotChange(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetTrotChange()
}

func (ies *ImpactEffectRPCService) CalEnergyAtmosphere(req *impactEffect.CalEnergyAtmosphereRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalEnergyAtmosphere(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetEnergyAtmosphere()
}

func (ies *ImpactEffectRPCService) CalEnergyBlastSurface(req *impactEffect.CalEnergyBlastSurfaceRequest) (float32, float32) {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalEnergyBlastSurface(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetEnergyBlast(), result.GetEnergySurface()
}

func (ies *ImpactEffectRPCService) CalMassOfWater(req *impactEffect.CalMassOfWaterRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalMassOfWater(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetMwater()
}

func (ies *ImpactEffectRPCService) CalVelocityProjectile(req *impactEffect.CalVelocityProjectileRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalVelocityProjectile(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetVseafloor()
}

func (ies *ImpactEffectRPCService) CalEnergyAtSeafloor(req *impactEffect.CalEnergyAtSeafloorRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalEnergyAtSeafloor(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetEnergySeafloor()
}

func (ies *ImpactEffectRPCService) CalEPIcentralAngle(req *impactEffect.CalEPIcentralAngleRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalEPIcentralAngle(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetDelta()
}

func (ies *ImpactEffectRPCService) CalScalingDiameterConstant(req *impactEffect.CalScalingDiameterConstantRequest) (float32, float32) {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalScalingDiameterConstant(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetCd(), result.GetBeta()
}

func (ies *ImpactEffectRPCService) CalAnglefac(req *impactEffect.CalAnglefacRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalAnglefac(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetAnglefac()
}

func (ies *ImpactEffectRPCService) CalWdiameter(req *impactEffect.CalWdiameterRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalWdiameter(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetWdiameter()
}

func (ies *ImpactEffectRPCService) CalTransientCraterDiameter(req *impactEffect.CalTransientCraterDiameterRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalTransientCraterDiameter(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetDtr()
}

func (ies *ImpactEffectRPCService) CalDepthr(req *impactEffect.CalDepthrRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalDepthr(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetDepthr()
}

func (ies *ImpactEffectRPCService) CalCdiamater(req *impactEffect.CalCdiamaterRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalCdiamater(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetCdiameter()
}

func (ies *ImpactEffectRPCService) CalDepthfr(req *impactEffect.CalDepthfrRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalDepthfr(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetDepthfr()
}

func (ies *ImpactEffectRPCService) CalVCrater(req *impactEffect.CalVCraterRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalVCrater(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetVCrater()
}

func (ies *ImpactEffectRPCService) CalVratio(req *impactEffect.CalVratioRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalVratio(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetVratio()
}

func (ies *ImpactEffectRPCService) CalVCraterVRation(req *impactEffect.CalVCraterVRationRequest) (float32, float32) {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalVCraterVRation(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetVCrater(), result.GetVratio()
}

func (ies *ImpactEffectRPCService) CalVMelt(req *impactEffect.CalVMeltRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalVMelt(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetVMelt()
}

func (ies *ImpactEffectRPCService) CalMratioAndMcratio(req *impactEffect.CalMratioAndMcratioRequest) (float32, float32) {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalMratioAndMcratio(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetMratio(), result.GetMcratio()
}

func (ies *ImpactEffectRPCService) CalEjectArrival(req *impactEffect.CalEjectArrivalRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalEjectArrival(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetEjectaArrival()
}

func (ies *ImpactEffectRPCService) CalEjectaThickness(req *impactEffect.CalEjectaThicknessRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalEjectaThickness(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetEjectaThickness()
}

func (ies *ImpactEffectRPCService) CalDFrag(req *impactEffect.CalDFragRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalDFrag(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetDFrag()
}

func (ies *ImpactEffectRPCService) CalThemal(req *impactEffect.CalThemalRequest) (
	float32, float32, float32, float32, float32, float32, float32, float32) {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalThemal(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetH(), result.GetRf(), result.GetThermalExposure(),
		result.GetNoRadiation(), result.GetMaxRadTime(), result.GetIrradiationTime(),
		result.GetMegatonFactor(), result.GetThermalPower()
}

func (ies *ImpactEffectRPCService) CalMagnitude(req *impactEffect.CalMagnitudeRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalMagnitude(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetMagnitude()
}

func (ies *ImpactEffectRPCService) CalMagnitude2(req *impactEffect.CalMagnitude2Request) (float32, float32) {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalMagnitude2(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetEffMag(), result.GetSeismicArrival()
}

func (ies *ImpactEffectRPCService) CalShockArrival(req *impactEffect.CalShockArrivalRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalShockArrival(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetShockArrival()
}

func (ies *ImpactEffectRPCService) CalVmax(req *impactEffect.CalVmaxRequest) (float32, float32) {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalVmax(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetVmax(), result.GetOpressure()
}

func (ies *ImpactEffectRPCService) CalShockDamage(req *impactEffect.CalShockDamageRequest) string {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalShockDamage(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetShockDamage()
}

func (ies *ImpactEffectRPCService) CalDecLevel(req *impactEffect.CalDecLevelRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.CalDecLevel(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetDecLevel()
}

func (ies *ImpactEffectRPCService) Cal_TsunamiArrivalTime(req *impactEffect.Cal_TsunamiArrivalTimeRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.Cal_TsunamiArrivalTime(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetTsunamiArrivalTime()
}

func (ies *ImpactEffectRPCService) Cal_WaveAmplitudeUpperLimit(req *impactEffect.Cal_WaveAmplitudeUpperLimitRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.Cal_WaveAmplitudeUpperLimit(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetWaveAmplitudeUpperLimit()
}

func (ies *ImpactEffectRPCService) Cal_WaveAmplitudeLowerLimit(req *impactEffect.Cal_WaveAmplitudeLowerLimitRequest) float32 {
	log.Printf("Getting (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ies.client.Cal_WaveAmplitudeLowerLimit(ctx, req)
	if err != nil {
		log.Fatalf("ies.client.Getresult failed: %v", err)
	}
	log.Println(result)

	return result.GetWaveAmplitudeLowerLimit()
}
