import { mergeMetrics, getPath, IPath } from "./src/merge";

const mergeFile = async (): Promise<string> => {
  console.log("Starting getting path...");
  const path: IPath = getPath();

  console.log("Starting formatting...");
  const updateFile = await mergeMetrics(path);

  console.log("Updated file:", updateFile);
  console.log("Done!");
  return updateFile;
};

const formatFile = async (path: string): Promise<string> => {
  console.log("Starting converting...");
  const updateFile = await convertUnit(path);
  return updateFile;
};

const main = async () => {
  // read cli arguments
  const cli = process.argv;
  const fn = cli[2];

  if (!fn) {
    throw new Error("Please choose function to run [merge | format | all]");
  }

  if (fn === "merge") {
    await mergeFile();
  } else if (fn === "format") {
    const path = cli[3];
    await formatFile(path);
  }
};

main();
