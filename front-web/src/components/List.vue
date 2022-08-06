<template>
  <div id="main" style="width:100%;">
    <div>
      <span>Please note: the results below are estimates based on current (limited) understanding of the impact process
        and come with large uncertainties; they should be used with caution, particularly in the case of peculiar input
        parameters. All values are given to three significant figures but this does not reflect the precision of the
        estimate. For more information about the uncertainty associated with our calculations and a full discussion of
        this program, please refer to this <a
          href="https://impact.ese.ic.ac.uk/ImpactEarth/ImpactEffects/effects.pdf">article</a> </span>
      <el-divider></el-divider>
    </div>
    <div>
      <div id="left" style="width:50%;float:left;">
        <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="150px" class="demo-ruleForm">
          <!-- Impactor Diameter -->
          <el-form-item label="Distance from Impact" required>
            <el-col :span=1>
              <el-tooltip>
                <div slot="content">Distance from Impact, Distance from Impact, Distance from Impact, Distance from
                  Impact,
                  Distance from Impact, Distance from Impact, Distance from Impact, Distance from Impact</div>
                <i class="el-icon-question icon-color"></i>
              </el-tooltip>
            </el-col>
            <el-col :span=13>
              <el-form-item prop="DistancefromImpact">
                <el-input v-model.number="ruleForm.DistancefromImpact"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span=7 :offset="1">
              <el-select v-model="ruleForm.DistancefromImpactA" placeholder="units">
                <el-option label="km" value="km"></el-option>
                <el-option label="miles" value="miles"></el-option>
              </el-select>
            </el-col>
          </el-form-item>
          <el-form-item label="Projectile Diameter" required>
            <el-col :span=1>
              <el-tooltip>
                <div slot="content">Impactor Diameter is a. Impactor Diameter is a. Impactor Diameter is a. Impactor
                  Diameter is a. Impactor Diameter is a.</div>
                <i class="el-icon-question icon-color"></i>
              </el-tooltip>
            </el-col>
            <el-col :span=13>
              <el-form-item prop="ImpactorDiameter">
                <el-input v-model.number="ruleForm.ImpactorDiameter"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span=7 :offset="1">
              <el-select v-model="ruleForm.ImpactorDiameterA" placeholder="units">
                <el-option label="miles" value="miles"></el-option>
                <el-option label="km" value="km"></el-option>
                <el-option label="feets" value="feets"></el-option>
              </el-select>
            </el-col>
          </el-form-item>

          <el-form-item label="Projectile Density" required>
            <el-col :span=1>
              <el-tooltip>
                <div slot="content">Impactor Density Impactor Density Impactor Density Impactor Density Impactor Density
                  Impactor Density Impactor Density</div>
                <i class="el-icon-question icon-color"></i>
              </el-tooltip>
            </el-col>
            <el-col :span=13>
              <el-form-item prop="ImpactorDensity">
                <el-input v-model.number="ruleForm.ImpactorDensity"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span=7 :offset="1">
              <el-select v-model="ruleForm.ImpactorDensityA" placeholder="units">
                <el-option label="kg/m^3" value="kg/m^3"></el-option>
              </el-select>
            </el-col>
          </el-form-item>

          <el-form-item label="Impactor Velocity" required>
            <el-col :span=1>
              <el-tooltip>
                <div slot="content">This is the velocity of the projectile before it enters the atmosphere. The minimum
                  impact velocity on Earth is 11 km/s. Typical impact velocities are 17 km/s for asteroids and 51 km/s
                  for
                  comets. The maximum Earth impact velocity for objects orbiting the sun is 72 km/s.</div>
                <i class="el-icon-question icon-color"></i>
              </el-tooltip>
            </el-col>
            <el-col :span=13>
              <el-form-item prop="ImpactorVelocity">
                <el-input v-model.number="ruleForm.ImpactorVelocity"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span=7 :offset=1>
              <el-select v-model="ruleForm.ImpactorVelocityA" placeholder="units">
                <el-option label="km/s" value="km/s"></el-option>
                <el-option label="m/s" value="m/s"></el-option>
              </el-select>
            </el-col>
          </el-form-item>

          <el-form-item label="Impactor Angle" required>
            <el-col :span=1>
              <el-tooltip>
                <div slot="content">Impactor Diameter is a. Impactor Diameter is a. Impactor Diameter is a. Impactor
                  Diameter is a. Impactor Diameter is a.</div>
                <i class="el-icon-question icon-color"></i>
              </el-tooltip>
            </el-col>
            <el-col :span=13>
              <el-form-item prop="ImpactorAngle">
                <el-input v-model.number="ruleForm.ImpactorAngle"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span=7 :offset=1>
              <el-select v-model="ruleForm.ImpactorAngleA" placeholder="units">
                <el-option label="degrees" value="degrees"></el-option>
              </el-select>
            </el-col>
          </el-form-item>

          <el-form-item label="Target Density" prop="TargerDensity">
            <el-radio-group v-model="ruleForm.TargerDensity">
              <el-row>
                <el-col :span="10" size="large">
                  <el-radio label="Water (meter)"></el-radio>
                </el-col>
                <el-col :span="14">
                  <el-form-item prop="WaterDepth" size="small">
                    <el-input v-model.number="ruleForm.WaterDepth"></el-input>
                  </el-form-item>
                </el-col>
              </el-row>
              <el-row>
                <el-radio label="Sedimentary Rock"></el-radio>
                <el-radio label="Crystalline Rock"></el-radio>
              </el-row>
            </el-radio-group>
          </el-form-item>

          <el-form-item>
            <el-button type="primary" @click="submitForm('ruleForm')">Calulate</el-button>
            <el-button @click="resetForm('ruleForm')">Reset</el-button>
          </el-form-item>
        </el-form>
      </div>

      <div id="mid" style="width:5%;float:left;">
        <el-divider direction="vertical"></el-divider>
      </div>
      <div id="right" style="width:45%;float:left;">
        <el-collapse v-model="activeNames" @change="handleChange">
          <el-collapse-item title="Energy" name="1">
            <div v-for="dis in Energydiscription">{{ dis }}</div>
          </el-collapse-item>
          <el-collapse-item title="Major Global Changes" name="2">
            <div v-for="dis in GlobalChangesdiscription">{{ dis }}</div>
          </el-collapse-item>
          <el-collapse-item v-if="AtmosphericShow" title="Atmospheric Entry:" name="2">
            <div v-for="dis in Atmosphericdiscription">{{ dis }}</div>
          </el-collapse-item>
          <el-collapse-item v-if="craterShow" title="Crater Dimensions" name="3">
            <div v-for="dis in Craterdiscription">{{ dis }}</div>
          </el-collapse-item>
          <el-collapse-item v-if="ejectaShow" title="Ejecta" name="4">
            <div v-for="dis in Ejectadiscription">{{ dis }}</div>
          </el-collapse-item>
          <el-collapse-item v-if="thermalShow" title="Thermal Radiation" name="4">
            <div v-for="dis in Themaldiscription">{{ dis }}</div>
          </el-collapse-item>
          <el-collapse-item v-if="seismicShow" title="Seismic Effects" name="4">
            <div v-for="dis in Seismicdiscription">{{ dis }}</div>
          </el-collapse-item>
          <el-collapse-item v-if="AirBlastShow" title="Air Blast" name="4">
            <a href="/">What does this mean?</a>
            <div v-for="dis in Airdiscription">{{ dis }}</div>
          </el-collapse-item>
          <el-collapse-item v-if="TsunamiShow" title="Tsunami Wave" name="6">
            <a href="/">What does this mean?</a>
            <div v-for="dis in Tsunamidiscription">{{ dis }}</div>
          </el-collapse-item>
        </el-collapse>
      </div>
    </div>
  </div>
</template>
<script>
import { ImpactService } from "@/models/List.js";
export default {
  data() {
    return {
      activeNames: ['1'],
      discription: "default",
      AtmosphericShow: false,
      craterShow: false,
      ejectaShow: false,
      thermalShow: false,
      seismicShow: false,
      AirBlastShow: false,
      TsunamiShow: false,
      Energydiscription: [''],
      GlobalChangesdiscription: [''],
      Atmosphericdiscription: [''],
      Craterdiscription: [''],
      Ejectadiscription: [''],
      Themaldiscription: [''],
      Seismicdiscription: [''],
      Airdiscription: [''],
      Tsunamidiscription: [''],
      ruleForm: {
        DistancefromImpact: '',
        DistancefromImpactA: 'miles',
        ImpactorDiameter: '',
        ImpactorDiameterA: 'miles',
        ImpactorDensity: '',
        ImpactorDensityA: 'kg/m^3',
        ImpactorAngle: '',
        ImpactorAngleA: 'degrees',
        ImpactorVelocity: '',
        ImpactorVelocityA: 'm/s',
        TargerDensity: 'Sedimentary Rock',
        WaterDepth: 0
      },
      rules: {
        DistancefromImpact: [
          { required: true, message: 'DistancefromImpact should not be empty!' },
          { type: 'number', message: 'DistancefromImpact should be numeric!' },
          { min: 1, max: 10000, type: 'number', message: 'DistancefromImpact range is [1, 10000]!' }

        ],
        DistancefromImpactA: [
          { required: true, message: 'please select units', trigger: 'change' }
        ],
        ImpactorDiameter: [
          { required: true, message: 'ImpactorDiameter should not be empty!' },
          { type: 'number', message: 'ImpactorDiameter should be numeric!' },
          { min: 1, max: 10000, type: 'number', message: 'ImpactorDiameter range is [1, 10000]!' }
        ],
        ImpactorDiameterA: [
          { required: true, message: 'please select units', trigger: 'change' }
        ],
        ImpactorDensity: [
          { required: true, message: 'ImpactorDensity should not be empty!' },
          { type: 'number', message: 'ImpactorDensity should be numeric!' },
          { min: 1, max: 10000000, type: 'number', message: 'ImpactorDensity range is [1, 10000000]!' },
        ],
        ImpactorDensityA: [
          { required: true, message: 'please select units', trigger: 'change' }
        ],
        ImpactorAngle: [
          { required: true, message: 'ImpactorAngle should not be empty!' },
          { type: 'number', message: 'ImpactorAngle should be numeric!' },
          { min: 0, max: 90, type: 'number', message: 'ImpactorAngle range is [0, 90]!' },
        ],
        ImpactorAngleA: [
          { required: true, message: 'please select units', trigger: 'change' }
        ],
        ImpactorVelocity: [
          { required: true, message: 'ImpactorVelocity should not be empty!' },
          { type: 'number', message: 'ImpactorVelocity should be numeric!' },
          { min: 0, max: 299999, type: 'number', message: 'ImpactorAngle range is [0, 299999]!' }
        ],
        ImpactorVelocityA: [
          { required: true, message: 'please select units', trigger: 'change' }
        ],
        TargerDensity: [
          { required: true, message: 'please select the target type', trigger: 'change' }
        ],
        WaterDepth: [
          { type: 'number', message: 'Water Depth should be numeric!' },
          { type: 'number', message: 'WaterDepth should be numeric!' },
          { min: 0, max: 10000, type: 'number', message: 'WaterDepth range is [1, 10000]!' },
        ]
      }
    };
  },
  methods: {
    submitForm(formName) {
      console.log(this.ruleForm.TargerDensity)

      this.$refs[formName].validate((valid) => {
        if (valid) {
          alert('Calculate!');
          // pack the parameters 
          var target_distance = Number(this.ruleForm.DistancefromImpact);
          if (this.ruleForm.DistancefromImpactA == 'km') {
            target_distance *= 1000;
          }
          var impactor_density = Number(this.ruleForm.ImpactorDensity);
          var impactor_diameter = Number(this.ruleForm.ImpactorDiameter);
          if (this.ruleForm.ImpactorDiameterA == 'km') {
            impactor_diameter *= 1000;
          }
          var impactor_velocity = Number(this.ruleForm.ImpactorVelocity);
          if (this.ruleForm.ImpactorVelocityA == 'km/s') {
            impactor_velocity *= 1000;
          }
          var impactor_degree = Number(this.ruleForm.ImpactorAngle);

          var target_density = 2500;
          var target_depth = 0;
          if (this.resetForm.TargerDensity == "Crystalline Rock") {
            target_density = 2750;
          } else {
            target_depth = Number(this.ruleForm.WaterDepth)
            target_density = 1000;
          }

          ImpactService(impactor_density, impactor_diameter, impactor_velocity,
            impactor_degree, target_density, target_depth, target_distance)
            .then(data => {
              console.log("request");
              console.log(data.data);
              this.Energydiscription = (data.data.energy_disc + data.data.rec_disc).split("/n");
              this.GlobalChangesdiscription = data.data.change_disc.split("/n");

              this.AtmosphericShow = data.data.atmos_disc != '';
              this.Atmosphericdiscription = data.data.atmos_disc.split("/n");

              this.craterShow = data.data.crater_disc != '';
              this.Craterdiscription = data.data.crater_disc.split("/n");

              this.ejectaShow = data.data.ejecta_disc != '';
              this.Ejectadiscription = data.data.ejecta_disc.split("/n");

              this.thermalShow = data.data.themal_disc != '';
              this.Themaldiscription = data.data.themal_disc.split("/n");

              this.seismicShow = data.data.seismic_disc != '';
              this.Seismicdiscription = data.data.seismic_disc.split("/n");

              this.AirBlastShow = data.data.airblast_disc != '';
              this.Airdiscription = data.data.airblast_disc.split("/n");

              this.TsunamiShow = data.data.tsunami_disc != '';
              this.Tsunamidiscription = data.data.tsunami_disc.split("/n");
            })
        } else {
          console.log('error submit!!');
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    handleChange(val) {
      console.log(val);
    }
  }
}
</script>


<style>
.demo-ruleForm {
  width: 100%;
}
</style>
