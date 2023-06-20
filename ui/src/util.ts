const URL_PATTERN = /^(https?|ftp):\/\/[^\s/$.?#].[^\s]*$/;

export const validateURL = (url: string) => {
  return URL_PATTERN.test(url);
};
