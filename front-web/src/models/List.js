import { ajax, ajaxPost } from "@/models/Service.js";

function ImpactService(impactor_density, impactor_diameter, impactor_velocity, 
    impactor_theta, target_density, target_depth, target_distance) {
    return ajaxPost("/simulator", {
        impactor_density: impactor_density,
        impactor_diameter: impactor_diameter,
        impactor_velocity: impactor_velocity,
        impactor_theta: impactor_theta, 
        target_density:target_density,
        target_depth: target_depth, 
        target_distance: target_distance
    });
  }

  export {ImpactService};