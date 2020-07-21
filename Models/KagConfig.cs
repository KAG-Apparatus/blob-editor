using System.Collections.Generic;
using System.Text.RegularExpressions;
using System.IO;
using System;
using System.Text;

namespace Blob_Editor
{
    public class KagConfig
    {
        private List<Element> elements;

        public KagConfig(List<Element> elements)
        {
            this.Elements = elements;
        }

        public List<Element> Elements { get => elements; set => elements = value; }
    }

    public class Parser
    {
        public static KagConfig Parse(string filepath)
        {
            List<Element> contents = new List<Element>();
            using (StreamReader reader = new StreamReader(filepath))
            {
                string line;
                int lineNumber = 0;

                while ((line = reader.ReadLine()) != null)
                {
                    lineNumber++;
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
                            } else
                            {
                                throw new Exception("Invalid CFG file");
                            }
                            break;
                    }
                }
            }

            return new KagConfig(contents);
        }

        public static LineType IdentifyLine(string line)
        {
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
    }
}