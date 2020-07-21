using System.Collections.Generic;
using System.Text.RegularExpressions;

namespace Blob_Editor
{
	public class KagConfig
	{
		Dictionary<string,string> keyValuePairs = new Dictionary<string, string>();

		KagConfig(string configRaw)
		{
			keyValuePairs = parse(configRaw);
		}

		static Dictionary<string,string> parse(string configRaw)
		{
			Dictionary<string,string> keyValuePairs = new Dictionary<string, string>();

			//Clean config file of whitespace.
			string configNoWhiteSpace = Regex.Replace(configRaw,"( {2,})|(\t{2,})", "");

			return keyValuePairs;
		}
	}
}