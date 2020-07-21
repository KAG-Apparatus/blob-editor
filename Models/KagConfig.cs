using System.Collections.Generic;
using System.Text.RegularExpressions;
using System.IO;
using System;

namespace Blob_Editor
{
    public class KagConfig
    {
        private List<KeyValuePair<string, List<string>>> contents;

        public KagConfig(List<KeyValuePair<string, List<string>>> contents)
        {
            this.contents = contents;
        }
    }

    public class Parser
    {
        public static KagConfig Parse(string filepath)
        {
            Stack<KeyValuePair<string, List<string>>> stackedContents = new Stack<KeyValuePair<string, List<string>>>();
            string pattern = @"(\@*\$*\w+)\ +=\ +(.*)";
            Regex regex = new Regex(pattern, RegexOptions.IgnoreCase);
            using (StreamReader reader = new StreamReader(filepath))
            {
                KeyValuePair<string, List<string>> entry;

                string line;
                int lineNumber = 0;

                while ((line = reader.ReadLine()) != null)
                {
                    lineNumber++;

                    // Empty line
                    if (line == "")
                    {
                        entry = new KeyValuePair<string, List<string>>(
                            "VoidInLine" + lineNumber,
                            new List<string>(new string[] { "" })
                        );
                        stackedContents.Push(entry);
                        continue;
                    }

                    // Comment line
                    if (line[0] == '#')
                    {
                        entry = new KeyValuePair<string, List<string>>(
                            "CommentInLine" + lineNumber,
                            new List<string>(new string[] { line })
                        );
                        stackedContents.Push(entry);
                        continue;
                    }

                    Match match = regex.Match(line);
                    if (match.Success)
                    {
                        string key = match.Groups[1].Captures[0].ToString();
                        string value = match.Groups[2].Captures[0].ToString();
                        entry = new KeyValuePair<string, List<string>>(
                            key,
                            new List<string>(new string[] { value })
                        );
                        stackedContents.Push(entry);
                        match = match.NextMatch();
                    }
                    // Empty entry
                    else if (line.Contains('='))
                    {
                        string key = line.Replace('=', '\0').Trim();
                        entry = new KeyValuePair<string, List<string>>(
                            key,
                            new List<string>(new string[] { "" })
                        );
                        stackedContents.Push(entry);
                    }
                    else
                    {
                        entry = stackedContents.Pop();
                        entry.Value.Add(line);
                        stackedContents.Push(entry);
                    }
                }
            }

            List<KeyValuePair<string, List<string>>> listedContents = new List<KeyValuePair<string, List<string>>>();
            foreach (KeyValuePair<string, List<string>> entry in stackedContents)
            {
                listedContents.Add(entry);
            }
            listedContents.Reverse();
            foreach (KeyValuePair<string, List<string>> entry in listedContents)
            {
                Console.WriteLine("{0} =", entry.Key);
                foreach (string value in entry.Value)
                {
                    Console.Write("\t{0}\n", value);
                }
            }
            return new KagConfig(listedContents);
        }
    }
}