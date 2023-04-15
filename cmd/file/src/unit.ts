import { UNIT_METADATA } from "./constant/unit-meta";
import { IPath } from "./interfaces/path";
import { streamStandardizeUnit } from "./stream/unit";

export const standardizeUnit = async (path: IPath): Promise<void> => {
  const { metricsWithOutputFile, metricsWithOutputAndStandardizedUnitFile } =
    path;

  await streamStandardizeUnit(
    metricsWithOutputFile,
    metricsWithOutputAndStandardizedUnitFile,
    UNIT_METADATA
  );
};
