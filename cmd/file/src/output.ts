import { IMetric } from "./interfaces/metric";
import { IPath } from "./interfaces/path";
import { streamAddTimeToFail, streamGetFailedResponse } from "./stream/output";

export const addTimeToFail = async (
  metrics: IMetric[],
  path: IPath
): Promise<void> => {
  const { logFile, metricsFile, metricsWithOutputFile } = path;
  const errorResponse: any[] = await streamGetFailedResponse(logFile);

  await streamAddTimeToFail(
    metricsFile,
    metricsWithOutputFile,
    errorResponse,
    metrics
  );
};
