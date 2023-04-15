import { getPath, addTimeToFail, mergeMetrics } from "./src";

const main = async () => {
  console.log("Starting getting path...");
  const path = getPath();

  console.log("Generating metrics file...");
  const generateFile = await mergeMetrics(path);

  console.log("Adding timestamp output to metrics file...");
  const updateFile = await addTimeToFail(path);

  console.log("Updated file:", updateFile);
  console.log("Done!");
};

main();
