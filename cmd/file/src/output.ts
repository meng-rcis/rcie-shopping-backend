import { IPath } from "./interfaces/path";
import { streamAddTimeToFail, streamGetFailedResponse } from "./stream/output";

export const addTimeToFail = async (path: IPath): Promise<string> => {
  const { logFile, metricsFile } = path;
  const errorResponse: any[] = await streamGetFailedResponse(logFile);
  await streamAddTimeToFail(metricsFile, metricsFile, errorResponse);

  return metricsFile;
};
