import { getPath, addTimestamp, mergeMetrics } from "./src";

const main = async () => {
  console.log("Starting getting path...");
  const path = getPath();

  console.log("Generating metrics file...");
  const generateFile = await mergeMetrics(path);

  console.log("Adding timestamp output to metrics file...");
  const updateFile = await addTimestamp(path);

  console.log("Updated file:", updateFile);
  console.log("Done!");
};

main();
