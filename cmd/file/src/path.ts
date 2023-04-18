import fs from "fs";
import { IPath } from "./interfaces/path";

const BASE_PATH = "../../external/jmeter";
const LOG_PATH = `${BASE_PATH}/log`;
const METRICS_PATH = `${BASE_PATH}/grafana`;
const FILE_TYPE = ".csv";
const BASE_NAME = "rcie-api-";

const CPU = "cpu";
const MEMORY = "memory";
const NETWORK = "network";
const TPS = "tps";
const LATENCY = "latency";
const RESPONSE = "response";
const COMBINED = "metrics";
const COMBINED_WITH_OUTPUT = "metrics-with-output";
const COMBINED_WITH_OUTPUT_AND_STANDARDIZED_UNIT =
  "metrics-with-output-and-standardized-unit";

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
