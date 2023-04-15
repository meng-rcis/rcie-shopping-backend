import fs from "fs";
import { IPath } from "./interfaces/path";

const BASE_PATH = "../../external/jmeter";
const LOG_PATH = `${BASE_PATH}/log`;
const METRICS_PATH = `${BASE_PATH}/grafana`;
const FILE_TYPE = ".csv";
const BASE_NAME = "rcie-api-";

export const getPath = (): IPath => {
  const cli = process.argv;
  const time = cli[2];
  const type = cli[3] || "default";

  if (!time) {
    throw new Error("Please enter the time of the test");
  }

  const isTypeValid =
    type.length === 3 && type.slice(0, 2) === "no" && isNumber(type.slice(2));
  if (type !== "default" && isTypeValid) {
    throw new Error("Please enter a valid type");
  }

  const name = `${time}_${BASE_NAME}${type}`;
  const logFile = `${LOG_PATH}/${name}${FILE_TYPE}`;
  const cpuFile = `${METRICS_PATH}/${name}/cpu${FILE_TYPE}`;
  const memoryFile = `${METRICS_PATH}/${name}/memory${FILE_TYPE}`;
  const networkFile = `${METRICS_PATH}/${name}/network${FILE_TYPE}`;
  const metricsFile = `${METRICS_PATH}/${name}/metrics${FILE_TYPE}`;
  const metricsWithOutputFile = `${METRICS_PATH}/${name}/metrics-with-output${FILE_TYPE}`;
  const metricsWithOutputAndStandardizedUnitFile = `${METRICS_PATH}/${name}/metrics-with-output-and-standardized-unit${FILE_TYPE}`;

  const isPathExist = verifyPath(logFile, cpuFile, memoryFile, networkFile);
  if (!isPathExist) {
    throw new Error("Please enter a valid path");
  }

  const path: IPath = {
    logFile,
    cpuFile,
    memoryFile,
    networkFile,
    metricsFile,
    metricsWithOutputFile,
    metricsWithOutputAndStandardizedUnitFile,
  };
  return path;
};

const isNumber = (value: string) => {
  return !isNaN(Number(value));
};

const verifyPath = (...paths: string[]): boolean => {
  paths.forEach((path) => {
    if (!fs.existsSync(path)) {
      return false;
    }
  });

  return true;
};
