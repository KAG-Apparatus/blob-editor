using System.Collections.Generic;
using System;
using System.Text.RegularExpressions;
using System.Text;
using System.IO;

namespace Blob_Editor
{
    public class CFG
    {
        private List<Element> elements;

        public CFG(List<Element> elements)
        {
            this.Elements = elements;
        }

        public CFG(string filepath)
        {
            List<Element> contents = new List<Element>();
            using (StreamReader reader = new StreamReader(filepath))
            {
                string line;
                while ((line = reader.ReadLine()) != null)
                {
                    LineType lineType = IdentifyLine(line);
                    switch (lineType)
                    {
                        case LineType.Comment:
                            contents.Add(new Comment(line));
                            break;
                        case LineType.RegularEntry:
                            contents.Add(new Entry(line));
                            break;
                        case LineType.EmptyEntry:
                            contents.Add(new Entry(line.Replace("=", "").Trim(), new List<string>()));
                            break;
                        case LineType.EmptyLine:
                            contents.Add(new Empty());
                            break;
                        case LineType.AvulseEntry:
                            if (contents.Count > 0)
                            {
                                Entry lastElement = (Entry)contents[contents.Count - 1];
                                lastElement.Append(line);
                                contents[contents.Count - 1] = (Element)lastElement;
                            }
                            else
                            {
                                throw new Exception("Invalid CFG file");
                            }
                            break;
                    }
                }
            }

            this.Elements = contents;
        }

        public static LineType IdentifyLine(string line)
        {
            line = line.Trim();

            if (line == "" || line == null)
            {
                return LineType.EmptyLine;
            }

            if (line[0] == '#')
            {
                return LineType.Comment;
            }

            if (!line.Contains('='))
            {
                return LineType.AvulseEntry;
            }

            if (line.Trim().EndsWith('='))
            {
                return LineType.EmptyEntry;
            }

            return LineType.RegularEntry;
        }

        public List<Element> Elements { get => elements; set => elements = value; }
    }

    public enum LineType
    {
        Comment,
        EmptyLine,
        EmptyEntry,
        RegularEntry,
        AvulseEntry
    }

    public interface Element
    {
        public string Key { get; set; }
        public List<string> ValueList { get; set; }
        string Print();
    }

    public class Comment : Element
    {
        private string key;

        bool HaveKey { get => true; }
        public string Key { get => this.key; set => this.key = value; }
        bool HaveValue { get => false; }
        public List<string> ValueList { get => null; set => throw new Exception("Can't set value to Comment"); }

        public Comment(string value)
        {
            this.key = value;
        }

        public string Print()
        {
            return this.key;
        }
    }

    public class Empty : Element
    {
        bool HaveKey { get => false; }
        public string Key { get => null; set => throw new Exception("Can't set key to Empty"); }
        bool HaveValue { get => false; }
        public List<string> ValueList { get => null; set => throw new Exception("Can't set value to Empty"); }

        public string Print()
        {
            return "";
        }
    }

    public class Entry : Element
    {
        private string key;
        private List<string> values;

        bool HaveKey { get => true; }
        public string Key { get => this.key; set => key = value; }
        bool HaveValue { get => true; }
        public List<string> ValueList { get => values; set => this.values = value; }

        public Entry(string key, List<string> values)
        {
            this.key = key;
            this.values = values;
        }

        public Entry(string line)
        {
            string pattern = @"(\@*\w*\ *\@*\$*\w+)\ *=\ *(.*)";
            Regex regex = new Regex(pattern, RegexOptions.IgnoreCase);
            Match match = regex.Match(line);
            if (match.Success)
            {
                this.key = match.Groups[1].Captures[0].ToString();
                string value = match.Groups[2].Captures[0].ToString();
                this.values = new List<string>(new string[] { value.Trim() });
            }
            else
            {
                throw new Exception("Invalid entry");

            }
        }

        public void Append(string value)
        {
            this.values.Add(value.Trim());
        }

        public string Print()
        {
            StringBuilder builder = new StringBuilder();
            builder.AppendFormat("{0} = ", this.key);
            builder.AppendJoin(',', this.values);
            return builder.ToString();
        }
    }
}