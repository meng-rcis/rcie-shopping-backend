import { getPath, addTimeToFail, mergeMetrics, generateMetrics } from "./src";
import { standardizeUnit } from "./src/unit";

const main = async () => {
  console.log("Starting getting path...");
  const path = getPath();

  console.log("Generating metrics file...");
  const metrics = generateMetrics(path);
  await mergeMetrics(metrics, path);

  console.log("Adding timestamp output to metrics file...");
  await addTimeToFail(metrics, path);

  console.log("Standardize unit...");
  await standardizeUnit(path);

  console.log("Done!");
};

main();
