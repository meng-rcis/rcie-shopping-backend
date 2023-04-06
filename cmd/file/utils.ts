export const findNearestFailedTime = (
  timeStamp: string,
  failedResponse: any[]
): string => {
  const timeStampInt = Number(timeStamp);
  let diff = -1;
  let skip = 0;
  let foundNearest = false;

  for (let i = 0; i < failedResponse.length; i++) {
    const error = failedResponse[i];
    const isTimeStampBigger = timeStampInt > Number(error.timeStamp);
    if (!isTimeStampBigger) {
      break;
    }
    skip++;
  }

  for (let i = skip; i < failedResponse.length; i++) {
    const error = failedResponse[i];
    const isTimeStampBigger = timeStampInt > Number(error.timeStamp);
    if (isTimeStampBigger) {
      break;
    }
    const currentDiff = Number(error.timeStamp) - timeStampInt;
    if (!foundNearest || currentDiff < diff) {
      diff = currentDiff;
      foundNearest = true;
    }
  }

  return foundNearest ? diff.toString() : "";
};
