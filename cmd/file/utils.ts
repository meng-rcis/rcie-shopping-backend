export const findNearestFailedTime = (
  timeStamp: string,
  failedResponse: any[]
): string => {
  const timeStampInt = Number(timeStamp);
  let skip = 0;

  for (let i = 0; i < failedResponse.length; i++) {
    const error = failedResponse[i];
    const isTimeStampBigger = timeStampInt > Number(error.timeStamp);
    if (!isTimeStampBigger) {
      break;
    }
    skip++;
  }

  const diff =
    skip < failedResponse.length
      ? Number(failedResponse[skip].timeStamp) - timeStampInt
      : "";
  return diff.toString();
};
