import fs from "fs";
import { IPath } from "./interface";

const BASE_PATH = "../../external/jmeter";
const LOG_PATH = `${BASE_PATH}/log`;
const METRICS_PATH = `${BASE_PATH}/grafana`;
const FILE_TYPE = ".csv";
const BASE_NAME = "rcie-api-";

export const getPath = (): IPath => {
  const cli = process.argv;
  const time = cli[3];
  const type = cli[4] || "default";

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
  const metricsFile = `${METRICS_PATH}/${name}/metrics${FILE_TYPE}`;

  const isPathExist = verifyPath(logFile, metricsFile);
  if (!isPathExist) {
    throw new Error("Please enter a valid path");
  }

  const path: IPath = { logFile, metricsFile };
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
