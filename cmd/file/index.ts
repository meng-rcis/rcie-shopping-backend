import { formatMetrics } from "./format";
import { IPath } from "./interface";
import { getPath } from "./path";

const main = async () => {
  console.log("Starting getting path...");
  const path: IPath = getPath();

  console.log("Starting formatting...");
  const updateFile = await formatMetrics(path);
};

main();
