import { getPath, addTimestamp } from "./src";

const main = async () => {
  console.log("Starting getting path...");
  const path = getPath();

  console.log("Adding timestamp output...");
  const updateFile = await addTimestamp(path);

  console.log("Updated file:", updateFile);
  console.log("Done!");
};

main();
