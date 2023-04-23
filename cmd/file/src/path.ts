import fs from "fs";
import { IPath } from "./interfaces/path";
import {
  BASE_NAME,
  COMBINED,
  COMBINED_WITH_OUTPUT,
  COMBINED_WITH_OUTPUT_AND_STANDARDIZED_UNIT,
  CPU,
  FILE_TYPE,
  LATENCY,
  LOG_PATH,
  MEMORY,
  METRICS_PATH,
  NETWORK,
  RESPONSE,
  TPS,
} from "./constant/path";

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
  const cpuFile = `${METRICS_PATH}/${name}/${CPU}${FILE_TYPE}`;
  const memoryFile = `${METRICS_PATH}/${name}/${MEMORY}${FILE_TYPE}`;
  const networkFile = `${METRICS_PATH}/${name}/${NETWORK}${FILE_TYPE}`;
  const tpsFile = `${METRICS_PATH}/${name}/${TPS}${FILE_TYPE}`;
  const latencyFile = `${METRICS_PATH}/${name}/${LATENCY}${FILE_TYPE}`;
  const responseFile = `${METRICS_PATH}/${name}/${RESPONSE}${FILE_TYPE}`;
  const metricsFile = `${METRICS_PATH}/${name}/${COMBINED}${FILE_TYPE}`;
  const metricsWithOutputFile = `${METRICS_PATH}/${name}/${COMBINED_WITH_OUTPUT}${FILE_TYPE}`;
  const metricsWithOutputAndStandardizedUnitFile = `${METRICS_PATH}/${name}/${COMBINED_WITH_OUTPUT_AND_STANDARDIZED_UNIT}${FILE_TYPE}`;

  const isPathExist = verifyPath(
    logFile,
    cpuFile,
    memoryFile,
    networkFile,
    tpsFile,
    latencyFile,
    responseFile
  );
  if (!isPathExist) {
    throw new Error("Please enter a valid path");
  }

  const path: IPath = {
    logFile,
    cpuFile,
    memoryFile,
    networkFile,
    tpsFile,
    latencyFile,
    responseFile,
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
