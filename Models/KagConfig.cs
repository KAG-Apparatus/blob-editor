using System.Collections.Generic;
using System.Text.RegularExpressions;
using System.IO;
using System;

namespace Blob_Editor
{
    public class KagConfig
    {
        private Dictionary<string, string[]> keyValuePairs;

        public KagConfig(Dictionary<string, string[]> keyValuePairs)
        {
            this.keyValuePairs = keyValuePairs;
        }
    }

    public class Parser
    {
        public static KagConfig Parse(string filepath)
        {
            Dictionary<string, string[]> dictionary = new Dictionary<string, string[]>();
            string pattern = @"(\@*\$*\w+)\ +=\ +(.*)";
            Regex r = new Regex(pattern, RegexOptions.IgnoreCase);
            using (StreamReader reader = new StreamReader(filepath))
            {
                string line;
                while ((line = reader.ReadLine()) != null)
                {
                    Match m = r.Match(line);
                    while (m.Success)
                    {
						string key = m.Groups[1].Captures[0].ToString();
						string value = m.Groups[2].Captures[0].ToString();
                        Console.WriteLine("{0} = {1}", key, value);
						dictionary.Add(key, new string[]{value});
                        m = m.NextMatch();
                    }
                }
            }
            return new KagConfig(dictionary);
        }
    }
}