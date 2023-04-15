import { IMetric } from "./interfaces/metric";
import { IPath } from "./interfaces/path";
import { streamAddTimeToFail, streamGetFailedResponse } from "./stream/output";

export const addTimeToFail = async (
  path: IPath,
  metrics: IMetric[]
): Promise<void> => {
  const { logFile, metricsFile } = path;
  const newFile = metricsFile.replace(".csv", "-with-output.csv");
  const errorResponse: any[] = await streamGetFailedResponse(logFile);

  await streamAddTimeToFail(metricsFile, newFile, errorResponse, metrics);
};
